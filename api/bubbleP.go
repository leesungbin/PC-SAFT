package api

import (
	"math"
)

type InitInputFragment struct {
	Tc float64
	Pc float64
	Tb float64
	x  float64
}

type Result struct {
	P float64
	y []float64
}

func BublP_init(components []InitInputFragment, T float64) (res Result) {
	res.y = make([]float64, len(components))
	Psat := make([]float64, len(components))
	for i, c := range components {
		B := math.Log(c.Pc/1.013) / ((1 / c.Tb) - (1 / c.Tc))
		A := math.Log(1.013) + B/c.Tb
		Psat[i] = math.Exp(A - B/T)
		res.P += Psat[i] * c.x
	}
	for i, c := range components {
		res.y[i] = c.x * Psat[i] / res.P
	}
	return
}
