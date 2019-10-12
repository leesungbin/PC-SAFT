package api

import "math"

type BT_Input struct {
	P  float64
	x_ []float64
}

type BT_Result struct {
	T  float64
	y_ []float64
	// Volume of vapor
	V_V float64
	// Volume of liquid
	V_L float64
}

type ty_init struct {
	T  float64
	y_ []float64
	Bp []float64
}

func (components *Comps) BublT_init(P float64, x_ []float64) (res ty_init) {
	nc := len(components.data)
	Ap := make([]float64, nc)
	res.Bp = make([]float64, nc)
	Ps := make([]float64, nc)

	res.y_ = make([]float64, nc)
	res.T = 0.

	for i := 0; i < nc; i++ {
		B := (math.Log(components.data[i].Pc) - math.Log(1.013)) / (1./components.data[i].Tb - 1./components.data[i].Tc)
		A := math.Log(1.013) + B/components.data[i].Tb
		Tbp := B / (A - math.Log(P))
		Ap[i] = A
		res.Bp[i] = B
		res.T += x_[i] * Tbp
	}

	for i := 0; i < nc; i++ {
		Ps[i] = math.Exp(Ap[i] - res.Bp[i]/res.T)
		res.y_[i] = x_[i] * Ps[i] / P
	}
	return
}

func (components *Comps) BublT(in BT_Input) (res BT_Result) {
	nc := len(components.data)

	initRes := components.BublT_init(in.P, in.x_)
	T := initRes.T
	y_ := initRes.y_
	B := 0.
	for i := 0; i < nc; i++ {
		B += in.x_[i] * initRes.Bp[i]
	}

	// Volume of Vapor & Liquid
	var V_V, V_L float64
	maxit := 3000
	for i := 0; i < maxit; i++ {
		fvi_L := GetVolumeInput{in.P, T, in.x_, "L"}
		V_L, _ = components.GetVolume(fvi_L)
		phi_L, fug_L := components.Fugacity(NewtonInput{V_L, in.P, T, in.x_})

		fvi_V := GetVolumeInput{in.P, T, y_, "V"}
		V_V, _ = components.GetVolume(fvi_V)
		phi_V, fug_V := components.Fugacity(NewtonInput{V_V, in.P, T, y_})

		ynew := make([]float64, nc)
		sumy := 0.
		for j := 0; j < nc; j++ {
			ynew[j] = in.x_[j] * (phi_L[j] / phi_V[j])
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
