package api

import (
	"fmt"
	"math"
)

type CalculationInput struct {
	T  float64
	P  float64
	x_ []float64
	y_ []float64
}
type CalculationResult struct {
	P float64
	T float64
	V float64
	x []float64
	y []float64
}
type Result struct {
	P float64
	y []float64
}

func (components *Comps) BublP_init(x_ []float64, T float64) (res Result) {
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

func (components *Comps) BublP(in CalculationInput) (res CalculationResult) {
	maxit := 3000
	initRes := components.BublP_init(in.x_, in.T)
	P := initRes.P
	y_ := initRes.y
	for i := 0; i < maxit; i++ {
		fvi_L := FindVolumeInput{P, in.T, in.x_, "L"}
		V_L, _ := components.GetVolume(fvi_L)
		phi_L, fug_L := components.Fugacity(NewtonInput{V_L, P, in.T, in.x_})

		fvi_V := FindVolumeInput{P, in.T, y_, "V"}
		V_V, _ := components.GetVolume(fvi_V)
		phi_V, fug_V := components.Fugacity(NewtonInput{V_V, P, in.T, y_})
		fmt.Printf("@@@working\n")

		// adjust y composition
		nc := len(components.data)
		ynew := make([]float64, nc)
		sumy := 0.
		for j := 0; j < nc; j++ {
			ynew[j] = y_[j] * phi_L[j] / phi_V[j]
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
	}
	return CalculationResult{P: P, y: y_}
}
