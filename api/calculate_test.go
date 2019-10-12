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

var NNN_FindV_newton_Input = NewtonInput{0.00010914379164188678, Pressure, Temperature, Composition_NNN}
var NNN_FindV_newton_Output = 0.00011453205172139417

func Test_FindV_newton(t *testing.T) {
	Vres, err := NNN_ethane_nHexane_cyclohexane.FindV_newton(NNN_FindV_newton_Input)
	if err != nil {
		t.Errorf("FindV_newton err : %v", err)
	} else {
		if !PassWithAccuracy4(Vres, NNN_FindV_newton_Output) {
			t.Errorf("got %v, expected %v\n", Vres, NNN_FindV_newton_Output)
		}
	}
}

func Test_Peos_P(t *testing.T) {
	got := NNN_ethane_nHexane_cyclohexane.Peos_P(NNN_FindV_newton_Input)
	want := 216.21032034626552
	if !PassWithAccuracy4(got, want) {
		t.Errorf("got %v, expected %v\n", got, want)
	}
}

var NNN_Fugacity_newton_Input = NewtonInput{0.00011453205172139417, Pressure, Temperature, Composition_NNN}
var NNN_Fugacity_Output_phi = []float64{2.9332610259971834, 0.05301216759321054, 0.037263385443932213}
var NNN_Fugacity_Output_fug = []float64{10.5983798014261, 0.28731320598651555, 0.3365976537047761}

func Test_Fugacity(t *testing.T) {
	got_phi, got_fug := NNN_ethane_nHexane_cyclohexane.Fugacity(NNN_Fugacity_newton_Input)
	want_phi := NNN_Fugacity_Output_phi
	want_fug := NNN_Fugacity_Output_fug

	for i := 0; i < len(got_phi); i++ {
		if !PassWithAccuracy4(got_phi[i], want_phi[i]) {
			t.Errorf("phi : got %v, expected %v\n", got_phi, want_phi)
		}
		if !PassWithAccuracy4(got_fug[i], want_fug[i]) {
			t.Errorf("phi : got %v, expected %v\n", got_phi, want_fug)
		}
	}
}

var NNN_GetVolume_Input = FindVolumeInput{Pressure, Temperature, Composition_NNN, "L"}

func Test_GetVolume(t *testing.T) {
	got, err := NNN_ethane_nHexane_cyclohexane.GetVolume(NNN_GetVolume_Input)
	want := 0.00011453205172139417
	if err != nil || !PassWithAccuracy4(got, want) {
		t.Errorf("err: %v\ngot %v, expected %v\n", err, got, want)
	}
}
