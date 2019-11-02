package test

import (
	"math"
	"testing"

	. "github.com/leesungbin/PC-SAFT/api"
)

var want_NP_benzene_water = CrossAssociatedValues{
	E_AB: [][]float64{
		[]float64{0.0, 1353.35},
		[]float64{1353.35, 2706.7},
	},
	K_AB: [][]float64{
		[]float64{0.0, 0.0},
		[]float64{0.0, 0.08924},
	},
}

func Test_PrepareCrossParameter(t *testing.T) {
	input := Comps{}
	input.Data = []Component{Benzene, Water_polar}
	got := PrepareCrossParameter(input)

	for i, v := range got.E_AB {
		for j, w := range v {
			if !PassWithAccuracy4(w, want_NP_benzene_water.E_AB[i][j]) {
				t.Errorf("%.4f %.4f", w, want_NP_benzene_water.E_AB[i][j])
				t.Errorf("water & benzene eAB expected %v but got %v, erorr at eAB[%d][%d]", want_NP_benzene_water.E_AB, got.E_AB, i, j)
			}
		}
	}
	for i, v := range got.K_AB {
		for j, w := range v {
			if !PassWithAccuracy4(w, want_NP_benzene_water.K_AB[i][j]) {
				t.Errorf("%.4f %.4f", w, want_NP_benzene_water.K_AB[i][j])
				t.Errorf("water & benzene kAB expected %v but got %v, erorr at kAB[%d][%d]", want_NP_benzene_water.K_AB, got.K_AB, i, j)
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

// 0.000...1 : 0의 갯수 : N개( N >= 1)
func PassWithAccuracyN(N int, compare, want float64) bool {
	if want == 0 {
		return math.Abs(compare-want) < 1e-10
	}
	res := math.Abs((compare - want) / want)
	if res < Pow(0.1, float64(N+1)) {
		return true
	}
	return false
}

var NNN_FindV_newton_Input = NewtonInput{V: 0.00010914379164188678, P: Pressure, T: Temperature, Z_: Composition_NNN}
var NNN_FindV_newton_Output = 0.00011453205172139417

func Test_FindV_newton(t *testing.T) {
	Vres, err := FindV_newton(NNN_ethane_nHexane_cyclohexane, NNN_FindV_newton_Input)
	if err != nil {
		t.Errorf("FindV_newton err : %v", err)
	} else {
		if !PassWithAccuracy4(Vres, NNN_FindV_newton_Output) {
			t.Errorf("got %v, expected %v\n", Vres, NNN_FindV_newton_Output)
		}
	}
}

func Test_Peos_P(t *testing.T) {
	got, _ := Peos_P(NNN_ethane_nHexane_cyclohexane, NNN_FindV_newton_Input)
	want := 216.21032034626552
	if !PassWithAccuracy4(got, want) {
		t.Errorf("got %v, expected %v\n", got, want)
	}
}

var NNN_Fugacity_newton_Input = NewtonInput{V: 0.00011453205172139417, P: Pressure, T: Temperature, Z_: Composition_NNN}
var NNN_Fugacity_Output_phi = []float64{2.9332610259971834, 0.05301216759321054, 0.037263385443932213}
var NNN_Fugacity_Output_fug = []float64{10.5983798014261, 0.28731320598651555, 0.3365976537047761}

func Test_Fugacity(t *testing.T) {
	got_phi, got_fug, _ := Fugacity(NNN_ethane_nHexane_cyclohexane, NNN_Fugacity_newton_Input)
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

var NNN_GetVolume_Input = GetVolumeInput{P: Pressure, T: Temperature, Z_: Composition_NNN, State: "L"}

func Test_GetVolume(t *testing.T) {
	got, err := GetVolume(NNN_ethane_nHexane_cyclohexane, NNN_GetVolume_Input)
	want := 0.00011453205172139417
	if err != nil || !PassWithAccuracy4(got, want) {
		t.Errorf("err: %v\ngot %v, expected %v\n", err, got, want)
	}
	// timeout error
	// got, err = NNN_ethane_nHexane_cyclohexane.GetVolume(GetVolumeInput{P: Pressure, T: Temperature, Z_: []float64{1, 0, 0}, State: "L"})
	// t.Errorf("err: %v\n", err)
}
