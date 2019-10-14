package test

import (
	"testing"

	. "github.com/leesungbin/PC-SAFT/api"
)

var want_BublP_init = PY_init{
	P: 18.065865443773628,
	Y: []float64{0.9669884021380635, 0.015248381227162324, 0.017763216634774145},
}

// Volume is not added to result
var want_BublP = BP_Result{P: 11.911044412248602, Y_: []float64{0.931106962189395, 0.031805774449269755, 0.037087479901090147}}
var want_BublP_NNP = BP_Result{P: 38.35158864951551, Y_: []float64{0.9531677842868514, 0.02934545182397382, 0.01748647899926908}}
var Composition_NNN = []float64{0.2, 0.3, 0.5}
var Composition_NNP = []float64{0.5, 0.3, 0.2}

var Temperature = 338.7
var Pressure = 18.065865443773628

func Test_BublP_init(t *testing.T) {
	got := NNN_ethane_nHexane_cyclohexane.BublP_init(Temperature, Composition_NNN)
	if !PassWithAccuracy4(got.P, want_BublP_init.P) {
		t.Errorf("Expected %v got %v", want_BublP_init.P, got.P)
	}
	for i, v := range got.Y {
		if !PassWithAccuracy4(v, want_BublP_init.Y[i]) {
			t.Errorf("Expected %v got %v", want_BublP_init.Y[i], v)
		}
	}
}

func Test_BublP(t *testing.T) {
	// start := time.Now()
	got, _ := NNN_ethane_nHexane_cyclohexane.BublP(BP_Input{T: Temperature, X_: Composition_NNN})
	// elapsed := time.Since(start)
	// t.Errorf("time required : %v\n", elapsed) // average 1.6~2 ms, python보다 10배 빠름
	// 오차 1% 미만
	if !PassWithAccuracyN(1, got.P, want_BublP.P) {
		t.Errorf("Expected %v got %v", want_BublP, got)
	}
	for i, v := range got.Y_ {
		if !PassWithAccuracyN(1, v, want_BublP.Y_[i]) {
			t.Errorf("Expected %v got %v", want_BublP, got)
		}
	}

	got, _ = NNP_ethane_nHexane_ethanol.BublP(BP_Input{T: Temperature, X_: Composition_NNP})
	if !PassWithAccuracyN(1, got.P, want_BublP_NNP.P) {
		t.Errorf("Expected %v got %v", want_BublP_NNP, got)
	}
	for i, v := range got.Y_ {
		if !PassWithAccuracyN(1, v, want_BublP_NNP.Y_[i]) {
			t.Errorf("Expected %v got %v", want_BublP_NNP, got)
		}
	}
}
