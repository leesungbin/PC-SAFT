package api

import (
	"testing"
)

var pcsaftInput = PCsaftInput{
	0.0001157378925614143, 338.7, []float64{0.1, 0.4, 0.5},
}
var polar_pcsaftInput = PCsaftInput{
	9.441696175765974e-5, 338.7, []float64{0.1, 0.4, 0.5},
}

// var pp_pcsaftInput = PCsaftInput{}

var want = &PCsaftResult{
	[]float64{0.3905800095221037, 0.01173591989723688, 0.006813729979993502}, 0.9531949926503014,
}
var polar_want = &PCsaftResult{
	[]float64{1.9682704680971759, 0.045308301346646374, 0.02237075117544578}, 0.14726927470808093,
}

// var pp_want = &PCsaftResult{}

func Test_pcsaft(t *testing.T) {
	got := TestComponents.PCsaft(pcsaftInput)
	if !PassWithAccuracy4(got.Z, want.Z) {
		t.Errorf("Z : got %v expected %v", got.Z, want.Z)
	}
	for i := 0; i < len(TestComponents.data); i++ {
		if !PassWithAccuracy4(got.Phi[i], want.Phi[i]) {
			t.Errorf("Phi : got %v expected %v", got.Phi[i], want.Phi[i])
		}
	}
	got = PolarTestComponents.PCsaft(polar_pcsaftInput)
	if !PassWithAccuracy4(got.Z, polar_want.Z) {
		t.Errorf("\npolar | Z : got %v expected %v", got.Z, polar_want.Z)
	}
	for i := 0; i < len(TestComponents.data); i++ {
		if !PassWithAccuracy4(got.Phi[i], polar_want.Phi[i]) {
			t.Errorf("\npolar | Phi : got %v expected %v", got.Phi[i], polar_want.Phi[i])
		}
	}
}
