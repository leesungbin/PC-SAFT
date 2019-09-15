package api

import (
	"math"
	"testing"
)

var Water_polar = Component{"water (polar)", 18.015, 647.3, 221.2, 0.344, 373.15, 1.0405, 2.9657, 175.15, 0.08924, 2706.7, 1.85, 0.66245}
var Benzene = Component{"benzene", 78.114, 562.2, 48.9, 0.212, 353.2, 2.4653, 3.6478, 287.35, 0.0, 0.0, 0.0, 0.0}

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

func Test_PrepareCrossParameter(t *testing.T) {
	input := Comps{}
	input.data = []Component{Benzene, Water_polar}
	got := PrepareCrossParameter(input)

	for i, v := range got.eAB {
		for j, w := range v {
			if !PassWithAccuracy4(w, benzene_water_res.eAB[i][j]) {
				t.Errorf("%.4f %.4f", w, benzene_water_res.eAB[i][j])
				t.Errorf("water & benzene eAB expected %v but got %v, erorr at eAB[%d][%d]", benzene_water_res.eAB, got.eAB, i, j)
			}
		}
	}
	for i, v := range got.kAB {
		for j, w := range v {
			if !PassWithAccuracy4(w, benzene_water_res.kAB[i][j]) {
				t.Errorf("%.4f %.4f", w, benzene_water_res.kAB[i][j])
				t.Errorf("water & benzene kAB expected %v but got %v, erorr at kAB[%d][%d]", benzene_water_res.kAB, got.kAB, i, j)
			}
		}
	}
}

// 유효숫자 4자리 검증
func PassWithAccuracy4(compare float64, want float64) bool {
	if want == 0 {
		return math.Abs(compare-want) < 1e-4
	}
	res := math.Abs((compare - want) / want)
	if res < 1e-4 {
		return true
	}
	return false
}
