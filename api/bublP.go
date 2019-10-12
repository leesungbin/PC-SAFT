package api

import (
	"fmt"
	"math"
)

type BP_Input struct {
	T  float64
	x_ []float64
}
type BP_Result struct {
	P   float64
	y_  []float64
	V_V float64
	V_L float64
}
type py_init struct {
	P float64
	y []float64
}

func (components *Comps) BublP_init(T float64, x_ []float64) (res py_init) {
	res.y = make([]float64, len(components.data))
	Psat := make([]float64, len(components.data))
	for i, c := range components.data {
		B := math.Log(c.Pc/1.013) / ((1 / c.Tb) - (1 / c.Tc))
		A := math.Log(1.013) + B/c.Tb
		Psat[i] = math.Exp(A - B/T)
		res.P += Psat[i] * x_[i]
	}
	for i, _ := range components.data {
		res.y[i] = x_[i] * Psat[i] / res.P
	}
	return
}

func (components *Comps) BublP(in BP_Input) (res BP_Result, err error) {
	var i int
	maxit := 3000
	initRes := components.BublP_init(in.T, in.x_)
	P := initRes.P
	y_ := initRes.y
	var V_V, V_L float64
	for i = 0; i < maxit; i++ {
		fvi_L := GetVolumeInput{P, in.T, in.x_, "L"}
		V_L, err = components.GetVolume(fvi_L)
		phi_L, fug_L := components.Fugacity(NewtonInput{V_L, P, in.T, in.x_})

		fvi_V := GetVolumeInput{P, in.T, y_, "V"}
		V_V, err = components.GetVolume(fvi_V)
		phi_V, fug_V := components.Fugacity(NewtonInput{V_V, P, in.T, y_})

		// adjust y composition
		nc := len(components.data)
		ynew := make([]float64, nc)
		sumy := 0.
		for j := 0; j < nc; j++ {
			ynew[j] = in.x_[j] * phi_L[j] / phi_V[j]
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
			break
		}
		P = Pnew
		y_ = ynew
		if math.Abs(V_V-V_L)/V_V < 1e-5 { // for single phase
			return BP_Result{P, y_, V_V, V_L}, nil
		}
	}
	fmt.Printf("bubbleP calculation iterated # : %d\n", i)
	return BP_Result{P, y_, V_V, V_L}, nil
}
