package api

import (
	"testing"
)

var water_polar = Component{
	Name: "water (polar)",
	Data: []float64{18.015, 647.3, 221.2, 0.344, 373.15, 1.0405, 2.9657, 175.15, 0.08924, 2706.7, 1.85, 0.66245},
}
var benzene = Component{
	Name: "benzene",
	Data: []float64{78.114, 562.2, 48.9, 0.212, 353.2, 2.4653, 3.6478, 287.35, 0.0, 0.0, 0.0, 0.0},
}

var benzene_water_res = CrossAssociatedValues{
	eAB: [][]float64{
		[]float64{0.0, 1353.35},
		[]float64{1353.35, 2706.7},
	},
	kAB: [][]float64{
		[]float64{0.0, 0.0},
		[]float64{0.0, 0.08924},
	},
}

func TestPrepareBinaryParameter(t *testing.T) {
	got := PrepareBinaryParameter(Components{benzene, water_polar})
	if !CalcSuccess999(got.eAB, benzene_water_res.eAB) {
		t.Errorf("water & benzene eAB expected %v but got %v", benzene_water_res.eAB, got.eAB)
	}
	if !CalcSuccess999(got.kAB, benzene_water_res.kAB) {
		t.Errorf("water & benzene kAB expected %v but got %v", benzene_water_res.kAB, got.kAB)
	}

	// [[0, 0, 0], [0, 0.0, 1353.35], [0, 1353.35, 2706.7]] [[0, 0, 0], [0, 0.0, 0.0], [0, 0.0, 0.08924]]
}

func CalcSuccess999(c1 [][]float64, c2 [][]float64) bool {
	for i := 0; i < len(c1[0]); i++ {
		cmp := c1[i][i] / c2[i][i]
		if cmp < 0.999 || cmp > 1.001 {
			return false
		}
	}
	return true
}
