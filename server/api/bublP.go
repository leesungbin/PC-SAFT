package api

import (
	"errors"
	"math"
)

type Eq_Input struct {
	T  float64   `json:"T"`
	P  float64   `json:"P"`
	X_ []float64 `json:"x"`
	Y_ []float64 `json:"y"`
}
type Eq_Result struct {
	P   float64   `json:"P"`
	T   float64   `json:"T"`
	X_  []float64 `json:"x"`
	Y_  []float64 `json:"y"`
	V_V float64   `json:"Vvap"`
	V_L float64   `json:"Vliq"`
}
type PY_init struct {
	P float64   `json:"P"`
	Y []float64 `json:"y"`
}

func BublP_init(components Comps, T float64, x_ []float64) (res PY_init) {
	res.Y = make([]float64, len(components.Data))
	Psat := make([]float64, len(components.Data))
	for i, c := range components.Data {
		B := math.Log(c.Pc/1.013) / ((1 / c.Tb) - (1 / c.Tc))
		A := math.Log(1.013) + B/c.Tb
		Psat[i] = math.Exp(A - B/T)
		res.P += Psat[i] * x_[i]
	}
	for i, _ := range components.Data {
		res.Y[i] = x_[i] * Psat[i] / res.P
	}
	return
}

func BublP(components Comps, in Eq_Input) (res Eq_Result, err error) {
	var i int
	var sumy float64
	maxit := 300
	initRes := BublP_init(components, in.T, in.X_)
	P := initRes.P
	y_ := initRes.Y
	var Vv_res, Vl_res float64
	for i = 0; i < maxit; i++ {
		fvi_L := GetVolumeInput{P, in.T, in.X_, "L"}
		V_L, err_l1 := GetVolume(components, fvi_L)
		if err_l1 != nil {
			return Eq_Result{}, err_l1
		}

		phi_L, fug_L, err_l2 := Fugacity(components, NewtonInput{V_L, P, in.T, in.X_})
		if err_l2 != nil {
			return Eq_Result{}, err_l2
		}
		// fmt.Printf("Liquid\nV_L : %v\nphi_L : %v\nfug_L : %v\n", V_L, phi_L, fug_L)

		fvi_V := GetVolumeInput{P, in.T, y_, "V"}
		V_V, err_v1 := GetVolume(components, fvi_V)
		if err_v1 != nil {
			return Eq_Result{}, err_v1
		}
		phi_V, fug_V, err_v2 := Fugacity(components, NewtonInput{V_V, P, in.T, y_})
		if err_v2 != nil {
			return Eq_Result{}, err_v2
		}
		// fmt.Printf("Vapor\nV_V : %v\nphi_V : %v\nfug_V : %v\n\n", V_V, phi_V, fug_V)
		// adjust y composition
		nc := len(components.Data)
		ynew := make([]float64, nc)
		sumy = 0.
		for j := 0; j < nc; j++ {
			ynew[j] = in.X_[j] * phi_L[j] / phi_V[j]
			sumy += ynew[j]
		}
		Pnew := P * sumy
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
		y_ = ynew
		if math.Abs(V_V-V_L)/V_V < 1e-5 { // for single phase
			return Eq_Result{P, in.T, in.X_, y_, V_V, V_L}, nil
		}
	}
	if sumy > 1.0001 {
		return Eq_Result{}, errors.New("calc failed : y > 1")
	}
	// fmt.Printf("bubbleP calculation iterated # : %d\n", i)
	return Eq_Result{P, in.T, in.X_, y_, Vv_res, Vl_res}, nil
}
