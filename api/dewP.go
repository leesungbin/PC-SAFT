package api

import "math"

type DP_Input struct {
	T  float64
	y_ []float64
}

type DP_Result struct {
	P   float64   `json:"P"`
	X_  []float64 `json:"x"`
	V_V float64   `json:"Vvap"`
	V_L float64   `json:"Vliq"`
}

type px_init struct {
	P  float64
	x_ []float64
}

func (components *Comps) DewP_init(T float64, y_ []float64) (res px_init) {
	nc := len(components.data)
	Ps := make([]float64, nc)
	res.x_ = make([]float64, nc)

	sum := 0.
	for i := 0; i < nc; i++ {
		B := (math.Log(components.data[i].Pc) - math.Log(1.013)) / (1./components.data[i].Tb - 1./components.data[i].Tc)
		A := math.Log(1.013) + B/components.data[i].Tb
		Ps[i] = math.Exp(A - B/T)
		sum += y_[i] / Ps[i]
	}
	res.P = 1. / sum

	for i := 0; i < nc; i++ {
		res.x_[i] = y_[i] * res.P / Ps[i]
	}
	return
}

func (components *Comps) DewP(in DP_Input) (res DP_Result) {
	nc := len(components.data)
	initRes := components.DewP_init(in.T, in.y_)
	P := initRes.P
	x_ := initRes.x_

	var V_V, V_L float64
	maxit := 3000
	for i := 0; i < maxit; i++ {
		gvi_L := GetVolumeInput{P, in.T, x_, "L"}
		V_L, _ = components.GetVolume(gvi_L)
		phi_L, fug_L := components.Fugacity(NewtonInput{V_L, P, in.T, x_})

		gvi_V := GetVolumeInput{P, in.T, in.y_, "V"}
		V_V, _ = components.GetVolume(gvi_V)
		phi_V, fug_V := components.Fugacity(NewtonInput{V_V, P, in.T, in.y_})

		xnew := make([]float64, nc)
		sumx := 0.

		for j := 0; j < nc; j++ {
			xnew[j] = in.y_[j] * (phi_V[j] / phi_L[j])
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
			break
		}
		P = Pnew
		x_ = xnew
		if math.Abs(V_V-V_L)/V_V < 1e-5 {
			return DP_Result{P, x_, V_V, V_L}
		}
	}
	return DP_Result{P, x_, V_V, V_L}
}
