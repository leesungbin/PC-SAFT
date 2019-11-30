package api

import (
	"errors"
	"math"
)

type TX_init struct {
	T  float64   `json:"T"`
	X_ []float64 `json:"x"`
	Bp []float64 `json:"B_p"`
}

func DewT_init(components Comps, P float64, y_ []float64) (res TX_init) {
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

func DewT(components Comps, in Eq_Input) (res Eq_Result, err error) {
	nc := len(components.Data)

	initRes := DewT_init(components, in.P, in.Y_)
	T := initRes.T
	x_ := initRes.X_

	B := 0.
	for i := 0; i < nc; i++ {
		B += in.Y_[i] * initRes.Bp[i]
	}

	var Vv_res, Vl_res float64
	maxit := 300
	for i := 0; i < maxit; i++ {
		gvi_L := GetVolumeInput{in.P, T, x_, "L"}
		V_L, err := GetVolume(components, gvi_L)
		if err != nil {
			return Eq_Result{}, err
		}
		phi_L, fug_L, _ := Fugacity(components, NewtonInput{V_L, in.P, T, x_})

		gvi_V := GetVolumeInput{in.P, T, in.Y_, "V"}
		V_V, err := GetVolume(components, gvi_V)
		if err != nil {
			return Eq_Result{}, err
		}
		phi_V, fug_V, err := Fugacity(components, NewtonInput{V_V, in.P, T, in.Y_})
		if err != nil {
			return Eq_Result{}, err
		}

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
			Vv_res = V_V
			Vl_res = V_L
			break
		}
		T += delT
		x_ = xnew

		if math.Abs(V_V-V_L)/V_V < 1e-5 {
			return Eq_Result{in.P, T, x_, in.Y_, V_V, V_L}, nil
		}
	}

	// 튀는 값 방지
	for x := range in.X_ {
		if x < 0 {
			return Eq_Result{}, errors.New("x < 0")
		}
	}
	for y := range in.Y_ {
		if y < 0 {
			return Eq_Result{}, errors.New("y < 0")
		}
	}
	return Eq_Result{in.P, T, x_, in.Y_, Vv_res, Vl_res}, nil
}
