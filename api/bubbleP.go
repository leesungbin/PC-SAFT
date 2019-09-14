package api

import (
	"math"
)

type InitInputFragment struct {
	component Component
	x         float64
}
type CalculationInput struct {
	T  float64
	P  float64
	x_ []float64
	y_ []float64
}
type Result struct {
	P float64
	y []float64
}

func BublP_init(components []InitInputFragment, T float64) (res Result) {
	res.y = make([]float64, len(components))
	Psat := make([]float64, len(components))
	for i, c := range components {
		B := math.Log(c.component.Pc/1.013) / ((1 / c.component.Tb) - (1 / c.component.Tc))
		A := math.Log(1.013) + B/c.component.Tb
		Psat[i] = math.Exp(A - B/T)
		res.P += Psat[i] * c.x
	}
	for i, c := range components {
		res.y[i] = c.x * Psat[i] / res.P
	}
	return
}

// func BublP(data CalculationInput)
