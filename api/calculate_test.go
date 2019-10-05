package api

import (
	"math"
	"testing"
)

var want_NP_benzene_water = CrossAssociatedValues{
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
			if !PassWithAccuracy4(w, want_NP_benzene_water.eAB[i][j]) {
				t.Errorf("%.4f %.4f", w, want_NP_benzene_water.eAB[i][j])
				t.Errorf("water & benzene eAB expected %v but got %v, erorr at eAB[%d][%d]", want_NP_benzene_water.eAB, got.eAB, i, j)
			}
		}
	}
	for i, v := range got.kAB {
		for j, w := range v {
			if !PassWithAccuracy4(w, want_NP_benzene_water.kAB[i][j]) {
				t.Errorf("%.4f %.4f", w, want_NP_benzene_water.kAB[i][j])
				t.Errorf("water & benzene kAB expected %v but got %v, erorr at kAB[%d][%d]", want_NP_benzene_water.kAB, got.kAB, i, j)
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

var findVnewtonInput = NewtonInput{0.0001157378925614143, 9.422949332244094, 338.7, []float64{0.1, 0.4, 0.5}}
var findVnewtonOutput = 0.00012126259852862196

var TestComponents_FindV_newton_Input = NewtonInput{0.0013702439054573965, 18.065865443773628, 338.7, []float64{0.2, 0.3, 0.5}}
var TestComponents_FindV_newton_Output = 0.0013851730855810051

func Test_FindV_newton(t *testing.T) {
	Vres, err := NNN_ethane_nHexane_cyclohexane.FindV_newton(TestComponents_FindV_newton_Input)
	if err != nil {
		t.Errorf("error panic : %v", err)
	} else {
		if !PassWithAccuracy4(Vres, TestComponents_FindV_newton_Output) {
			t.Errorf("got %v, expected %v\n", Vres, TestComponents_FindV_newton_Output)
		}
	}
}

var Peos_POutput = 222.49341754650348

func Test_Peos_P(t *testing.T) {
	f := NNN_ethane_nHexane_cyclohexane.Peos_P(findVnewtonInput)
	if !PassWithAccuracy4(f, Peos_POutput) {
		t.Errorf("got %v, expected %v\n", f, Peos_POutput)
	}
}
