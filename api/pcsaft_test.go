package api

import (
	"testing"
)

var testInput = PCsaftInput{
	0.0001157378925614143, 338.7, []float64{0.1, 0.4, 0.5}, []float64{0, 0, 0}, Components{Ethane, Nhexane, Cyclohexane},
}
var want = PCsaftResult{
	[]float64{3.5206e-10, 3.7982999999999997e-10, 3.8499e-10}, 0.9531949926503014,
}

func Test_pcsaft(t *testing.T) {
	got := PCsaft(testInput)
	if !PassWithAccuracy4(got.Z, want.Z) {
		t.Errorf("Z : got %v expected %v", got.Z, want.Z)
	}
	for i := 0; i < len(testInput.component); i++ {
		if !PassWithAccuracy4(got.Phi[i], want.Phi[i]) {
			t.Errorf("Phi : got %v expected %v", got.Phi[i], want.Phi[i])
		}
	}
}
