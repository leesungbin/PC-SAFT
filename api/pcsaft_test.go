package api

import (
	"testing"
)

var testInput = PCsaftInput{
	0.0001157378925614143, 338.7, []float64{0.1, 0.4, 0.5}, []float64{0, 0, 0}, Components{Ethane, Nhexane, Cyclohexane},
}
var want = PCsaftResult{
	1, 0.9531949926503014,
}

func Test_pcsaft(t *testing.T) {
	got := PCsaft(testInput)
	if !PassWithAccuracy4(got.Z, want.Z) {
		t.Errorf("got %v expected %v", got.Z, want.Z)
	}
}
