package api

import "math"

type BT_Input struct {
	P  float64   `json:"P"`
	X_ []float64 `json:"x"`
}

type BT_Result struct {
	T  float64   `json:"T"`
	Y_ []float64 `json:"y"`
	// Volume of vapor
	V_V float64 `json:"Vvap"`
	// Volume of liquid
	V_L float64 `json:"Vliq"`
}

type TY_init struct {
	T  float64   `json:"T"`
	Y_ []float64 `json:"y"`
	Bp []float64 `json:"B_p"`
}

func (components *Comps) BublT_init(P float64, x_ []float64) (res TY_init) {
	nc := len(components.Data)
	Ap := make([]float64, nc)
	res.Bp = make([]float64, nc)
	Ps := make([]float64, nc)

	res.Y_ = make([]float64, nc)
	res.T = 0.

	for i := 0; i < nc; i++ {
		B := (math.Log(components.Data[i].Pc) - math.Log(1.013)) / (1./components.Data[i].Tb - 1./components.Data[i].Tc)
		A := math.Log(1.013) + B/components.Data[i].Tb
		Tbp := B / (A - math.Log(P))
		Ap[i] = A
		res.Bp[i] = B
		res.T += x_[i] * Tbp
	}

	for i := 0; i < nc; i++ {
		Ps[i] = math.Exp(Ap[i] - res.Bp[i]/res.T)
		res.Y_[i] = x_[i] * Ps[i] / P
	}
	return
}

func (components *Comps) BublT(in BT_Input) (res BT_Result) {
	nc := len(components.Data)

	initRes := components.BublT_init(in.P, in.X_)
	T := initRes.T
	y_ := initRes.Y_
	B := 0.
	for i := 0; i < nc; i++ {
		B += in.X_[i] * initRes.Bp[i]
	}

	// Volume of Vapor & Liquid
	var V_V, V_L float64
	maxit := 3000
	for i := 0; i < maxit; i++ {
		fvi_L := GetVolumeInput{in.P, T, in.X_, "L"}
		V_L, _ = components.GetVolume(fvi_L)
		phi_L, fug_L, _ := components.Fugacity(NewtonInput{V_L, in.P, T, in.X_})

		fvi_V := GetVolumeInput{in.P, T, y_, "V"}
		V_V, _ = components.GetVolume(fvi_V)
		phi_V, fug_V, _ := components.Fugacity(NewtonInput{V_V, in.P, T, y_})

		ynew := make([]float64, nc)
		sumy := 0.
		for j := 0; j < nc; j++ {
			ynew[j] = in.X_[j] * (phi_L[j] / phi_V[j])
			sumy += ynew[j]
		}

		delT := -math.Log(sumy) * T * T / B
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
		y_ = ynew

		if math.Abs(V_V-V_L)/V_V < 1e-5 { // for single phase
			return BT_Result{T, y_, V_V, V_L}
		}
	}
	return BT_Result{T, y_, V_V, V_L}
}
