package api

import (
	"testing"
)

var testInput = PCsaftInput{
	0.0001157378925614143, 338.7, []float64{0.1, 0.4, 0.5},
}
var want = &PCsaftResult{
	[]float64{0.3905800095221037, 0.01173591989723688, 0.006813729979993502}, 0.9531949926503014,
}

func Test_pcsaft(t *testing.T) {
	got := TestComponents.PCsaft(testInput)
	// t.Errorf("testing")
	if !PassWithAccuracy4(got.Z, want.Z) {
		t.Errorf("Z : got %v expected %v", got.Z, want.Z)
	}
	for i := 0; i < len(TestComponents.data); i++ {
		if !PassWithAccuracy4(got.Phi[i], want.Phi[i]) {
			t.Errorf("Phi : got %v expected %v", got.Phi[i], want.Phi[i])
		}
	}
}
