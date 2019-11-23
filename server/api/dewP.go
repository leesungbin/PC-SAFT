package api

import "math"

type PX_init struct {
	P  float64   `json:"P"`
	X_ []float64 `json:"x"`
}

func DewP_init(components Comps, T float64, y_ []float64) (res PX_init) {
	nc := len(components.Data)
	Ps := make([]float64, nc)
	res.X_ = make([]float64, nc)

	sum := 0.
	for i := 0; i < nc; i++ {
		B := (math.Log(components.Data[i].Pc) - math.Log(1.013)) / (1./components.Data[i].Tb - 1./components.Data[i].Tc)
		A := math.Log(1.013) + B/components.Data[i].Tb
		Ps[i] = math.Exp(A - B/T)
		sum += y_[i] / Ps[i]
	}
	res.P = 1. / sum

	for i := 0; i < nc; i++ {
		res.X_[i] = y_[i] * res.P / Ps[i]
	}
	return
}

func DewP(components Comps, in Eq_Input) (res Eq_Result, err error) {
	nc := len(components.Data)
	initRes := DewP_init(components, in.T, in.Y_)
	P := initRes.P
	x_ := initRes.X_

	var Vv_res, Vl_res float64
	maxit := 300
	for i := 0; i < maxit; i++ {
		gvi_L := GetVolumeInput{P, in.T, x_, "L"}
		V_L, err_l1 := GetVolume(components, gvi_L)
		if err_l1 != nil {
			return Eq_Result{}, err_l1
		}

		phi_L, fug_L, err_l2 := Fugacity(components, NewtonInput{V_L, P, in.T, x_})
		if err_l2 != nil {
			return Eq_Result{}, err_l2
		}

		gvi_V := GetVolumeInput{P, in.T, in.Y_, "V"}
		V_V, err_v1 := GetVolume(components, gvi_V)
		if err_v1 != nil {
			return Eq_Result{}, err_v1
		}
		phi_V, fug_V, err_v2 := Fugacity(components, NewtonInput{V_V, P, in.T, in.Y_})
		if err_v1 != nil {
			return Eq_Result{}, err_v2
		}

		xnew := make([]float64, nc)
		sumx := 0.

		for j := 0; j < nc; j++ {
			xnew[j] = in.Y_[j] * (phi_V[j] / phi_L[j])
			sumx += xnew[j]
		}
		Pnew := P / sumx

		converged := true
		for j := 0; j < nc; j++ {
			if math.Abs(fug_L[j]-fug_V[j]) > 1e-5 {
				converged = false
			}
		}
		if math.Abs(Pnew-P) < 1e-5 && converged {
			Vv_res = V_V
			Vl_res = V_L
			break
		}
		P = Pnew
		x_ = xnew
		if math.Abs(V_V-V_L)/V_V < 1e-5 {
			return Eq_Result{P, in.T, x_, in.Y_, V_V, V_L}, nil
		}
	}
	return Eq_Result{P, in.T, x_, in.Y_, Vv_res, Vl_res}, nil
}
