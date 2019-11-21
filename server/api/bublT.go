package api

import (
	"errors"
	"fmt"
	"math"
)

type TY_init struct {
	T  float64   `json:"T"`
	Y_ []float64 `json:"y"`
	Bp []float64 `json:"B_p"`
}

func BublT_init(components Comps, P float64, x_ []float64) (res TY_init) {
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

func BublT(components Comps, in Eq_Input) (res Eq_Result, err error) {
	var sumy float64
	nc := len(components.Data)

	initRes := BublT_init(components, in.P, in.X_)
	T := initRes.T
	y_ := initRes.Y_
	B := 0.
	for i := 0; i < nc; i++ {
		B += in.X_[i] * initRes.Bp[i]
	}

	// Volume of Vapor & Liquid
	var Vv_res, Vl_res float64
	maxit := 3000
	for i := 0; i < maxit; i++ {
		fvi_L := GetVolumeInput{in.P, T, in.X_, "L"}
		V_L, err_l1 := GetVolume(components, fvi_L)
		// fmt.Printf("V_L: %v\nfvi_L: %v\n", V_L, fvi_L)
		if err_l1 != nil {
			return Eq_Result{}, err_l1
		}

		phi_L, fug_L, err_l2 := Fugacity(components, NewtonInput{V_L, in.P, T, in.X_})
		// fmt.Printf("phi_L: %v\n fug_L : %v\n", phi_L, fug_L)
		if err_l2 != nil {
			return Eq_Result{}, err_l2
		}

		fvi_V := GetVolumeInput{in.P, T, y_, "V"}
		V_V, err_v1 := GetVolume(components, fvi_V)
		// fmt.Printf("V_V: %v\nfvi_V: %v\n", V_V, fvi_V)
		if err_v1 != nil {
			return Eq_Result{}, err_v1
		}
		phi_V, fug_V, err_v2 := Fugacity(components, NewtonInput{V_V, in.P, T, y_})
		// fmt.Printf("phi_V: %v\n fug_V : %v\n", phi_V, fug_V)
		if err_v2 != nil {
			return Eq_Result{}, err_v2
		}

		ynew := make([]float64, nc)
		sumy = 0.
		for j := 0; j < nc; j++ {
			ynew[j] = in.X_[j] * (phi_L[j] / phi_V[j])
			sumy += ynew[j]
		}

		delT := -math.Log(sumy) * T * T / B
		// fmt.Printf("delT: %v\n\n", delT)
		converged := true
		for j := 0; j < nc; j++ {
			if math.Abs(fug_L[j]-fug_V[j]) > 1e-5 {
				converged = false
			}
		}
		if math.Abs(delT) < 1e-3 && converged {
			Vv_res = V_V
			Vl_res = V_L
			break
		}

		T += delT
		y_ = ynew

		if math.Abs(V_V-V_L)/V_V < 1e-5 { // for single phase
			return Eq_Result{in.P, T, in.X_, y_, V_V, V_L}, nil
		}
	}
	if sumy > 1.0001 {
		return Eq_Result{}, errors.New(fmt.Sprintf("calc failed : y > 1, %v", sumy))
	}
	// fmt.Printf("max iter, in.X : %v\n", in.X_)
	return Eq_Result{in.P, T, in.X_, y_, Vv_res, Vl_res}, nil
}
