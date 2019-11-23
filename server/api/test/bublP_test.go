package test

import (
	"fmt"
	"testing"
	"time"
	// "time"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

var want_BublP_init = PY_init{
	P: 18.065865443773628,
	Y: []float64{0.9669884021380635, 0.015248381227162324, 0.017763216634774145},
}

// Volume is not added to result
var want_BublP = Eq_Result{P: 11.911044412248602, Y_: []float64{0.931106962189395, 0.031805774449269755, 0.037087479901090147}}
var want_BublP_NNP = Eq_Result{P: 38.35158864951551, Y_: []float64{0.9531677842868514, 0.02934545182397382, 0.01748647899926908}}
var Composition_NNN = []float64{0.2, 0.3, 0.5}
var Composition_NNP = []float64{0.5, 0.3, 0.2}

var Temperature = 338.7
var Pressure = 18.065865443773628

func Test_BublP_init(t *testing.T) {
	got := BublP_init(NNN_ethane_nHexane_cyclohexane, Temperature, Composition_NNN)
	if !PassWithAccuracy4(got.P, want_BublP_init.P) {
		t.Errorf("Expected %v got %v", want_BublP_init.P, got.P)
	}
	for i, v := range got.Y {
		if !PassWithAccuracy4(v, want_BublP_init.Y[i]) {
			t.Errorf("Expected %v got %v", want_BublP_init.Y[i], v)
		}
	}
}

func Test_BublP_init_PPP(t *testing.T) {
	start := time.Now()
	got := BublP_init(PPP_methanol_water_aceticacid, 300, PPP_composition)
	elapsed := time.Since(start)
	fmt.Printf("time required : %v\n", elapsed) // 1.359 us
	P_want := 0.16185190274476668
	Y_want := []float64{0.935211528952177, 0.05633965101420472, 0.008448820033618156}
	if !PassWithAccuracy4(got.P, P_want) {
		t.Errorf("Expected %v got %v\n", P_want, got.P)
	}
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(got.Y[i], Y_want[i]) {
			t.Errorf("Expected %v got %v\n", Y_want[i], got.Y[i])
		}
	}
}

func Test_BublP(t *testing.T) {
	start := time.Now()
	got, _ := BublP(NNN_ethane_nHexane_cyclohexane, Eq_Input{T: Temperature, X_: Composition_NNN})
	elapsed := time.Since(start)
	fmt.Printf("time required : %v\n", elapsed) // average 1.6~2 ms, python보다 10배 빠름
	// 오차 1% 미만
	if !PassWithAccuracyN(2, got.P, want_BublP.P) {
		t.Errorf("Expected %v got %v", want_BublP, got)
	}
	for i, v := range got.Y_ {
		if !PassWithAccuracyN(2, v, want_BublP.Y_[i]) {
			t.Errorf("Expected %v got %v", want_BublP, got)
		}
	}

	got, _ = BublP(NNP_ethane_nHexane_ethanol, Eq_Input{T: Temperature, X_: Composition_NNP})
	if !PassWithAccuracyN(2, got.P, want_BublP_NNP.P) {
		t.Errorf("Expected %v got %v", want_BublP_NNP, got)
	}
	for i, v := range got.Y_ {
		if !PassWithAccuracyN(2, v, want_BublP_NNP.Y_[i]) {
			t.Errorf("Expected %v got %v", want_BublP_NNP, got)
		}
	}
}

func Test_BublP_PPP(t *testing.T) {
	res, err := BublP(PPP_methanol_water_aceticacid, Eq_Input{T: 300, X_: PPP_composition})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	P_want := 0.15725065375735317
	Y_want := []float64{0.8884315687198456, 0.09211096673754716, 0.01943383543373948}
	if !PassWithAccuracy4(res.P, P_want) {
		t.Errorf("Expected %v got %v\n", P_want, res.P)
	}
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(res.Y_[i], Y_want[i]) {
			t.Errorf("Expectd %v got %v\n", Y_want, res.Y_)
			break
		}
	}

}
