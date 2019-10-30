package api

import "math"

type DT_Input struct {
	P  float64   `json:"P"`
	Y_ []float64 `json:"y"`
}

type DT_Result struct {
	T   float64   `json:"T"`
	X_  []float64 `json:"x"`
	V_V float64   `json:"Vvap"`
	V_L float64   `json:"Vliq"`
}

type TX_init struct {
	T  float64   `json:"T"`
	X_ []float64 `json:"x"`
	Bp []float64 `json:"B_p"`
}

func (components *Comps) DewT_init(P float64, y_ []float64) (res TX_init) {
	nc := len(components.Data)
	Ap := make([]float64, nc)
	res.Bp = make([]float64, nc)
	Ps := make([]float64, nc)

	for i := 0; i < nc; i++ {
		B := (math.Log(components.Data[i].Pc) - math.Log(1.013)) / (1./components.Data[i].Tb - 1./components.Data[i].Tc)
		A := math.Log(1.013) + B/components.Data[i].Tb
		Tbp := B / (A - math.Log(P))
		Ap[i] = A
		res.Bp[i] = B

		res.T += y_[i] * Tbp
	}

	res.X_ = make([]float64, nc)
	for i := 0; i < nc; i++ {
		Ps[i] = math.Exp(Ap[i] - res.Bp[i]/res.T)
		res.X_[i] = y_[i] * P / Ps[i]
	}
	return
}

func (components *Comps) DewT(in DT_Input) (res DT_Result) {
	nc := len(components.Data)

	initRes := components.DewT_init(in.P, in.Y_)
	T := initRes.T
	x_ := initRes.X_

	B := 0.
	for i := 0; i < nc; i++ {
		B += in.Y_[i] * initRes.Bp[i]
	}

	var V_V, V_L float64
	maxit := 3000
	for i := 0; i < maxit; i++ {
		gvi_L := GetVolumeInput{in.P, T, x_, "L"}
		V_L, _ := components.GetVolume(gvi_L)
		phi_L, fug_L, _ := components.Fugacity(NewtonInput{V_L, in.P, T, x_})

		gvi_V := GetVolumeInput{in.P, T, in.Y_, "V"}
		V_V, _ := components.GetVolume(gvi_V)
		phi_V, fug_V, _ := components.Fugacity(NewtonInput{V_V, in.P, T, in.Y_})

		xnew := make([]float64, nc)
		sumx := 0.

		for j := 0; j < nc; j++ {
			xnew[j] = in.Y_[j] * (phi_V[j] / phi_L[j])
			sumx += xnew[j]
		}

		delT := math.Log(sumx) * T * T / B
		converged := true
		for j := 0; j < nc; j++ {
			if math.Abs(fug_L[j]-fug_V[j]) > 1e-5 {
				converged = false
			}
		}
		if math.Abs(delT) < 1e-3 && converged {
			break
		}
		T += delT
		x_ = xnew

		if math.Abs(V_V-V_L)/V_V < 1e-5 {
			return DT_Result{T, x_, V_V, V_L}
		}
	}
	return DT_Result{T, x_, V_V, V_L}
}
