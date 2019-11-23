package test

import (
	"testing"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

var input_PCsaft = PCsaftInput{
	V: 0.0001157378925614143, T: 338.7, X_: []float64{0.1, 0.4, 0.5},
}
var input_PCsaft_polar = PCsaftInput{
	V: 9.441696175765974e-5, T: 338.7, X_: []float64{0.1, 0.4, 0.5},
}

// var pp_pcsaftInput = PCsaftInput{}

var want = &PCsaftResult{
	Phi: []float64{0.3905800095221037, 0.01173591989723688, 0.006813729979993502}, Z: 0.9531949926503014,
}
var want_polar = &PCsaftResult{
	Phi: []float64{1.9682704680971759, 0.045308301346646374, 0.02237075117544578}, Z: 0.14726927470808093,
}

// var pp_want = &PCsaftResult{}

func Test_pcsaft(t *testing.T) {
	got, err := PCsaft(NNN_ethane_nHexane_cyclohexane, input_PCsaft)
	if err != nil {
		t.Errorf("%v\n", err)
	} else if !PassWithAccuracy4(got.Z, want.Z) {
		t.Errorf("Z : got %v expected %v", got.Z, want.Z)
	}
	for i := 0; i < len(NNN_ethane_nHexane_cyclohexane.Data); i++ {
		if !PassWithAccuracy4(got.Phi[i], want.Phi[i]) {
			t.Errorf("Phi : got %v expected %v", got.Phi[i], want.Phi[i])
		}
	}
	got, err = PCsaft(NNP_ethane_nHexane_ethanol, input_PCsaft_polar)
	if err != nil {
		t.Errorf("%v\n", err)
	} else if !PassWithAccuracy4(got.Z, want_polar.Z) {
		t.Errorf("\npolar | Z : got %v expected %v", got.Z, want_polar.Z)
	}
	for i := 0; i < len(NNN_ethane_nHexane_cyclohexane.Data); i++ {
		if !PassWithAccuracy4(got.Phi[i], want_polar.Phi[i]) {
			t.Errorf("\npolar | Phi : got %v expected %v", got.Phi[i], want_polar.Phi[i])
		}
	}
}

func Test_PPP_pcsaft(t *testing.T) {
	got, err := PCsaft(PPP_methanol_water_aceticacid, PCsaftInput{V: 0.00003512664042169545, T: 300, X_: PPP_composition})
	// t.Errorf("%v\n%v\n", err, got)
	Z_want := 0.8658783139415385
	if err != nil {
		t.Errorf("%v\n", err)
	} else if !PassWithAccuracy4(got.Z, Z_want) {
		t.Errorf("Z / got %v expected %v", got.Z, Z_want)
	}
	Phi_want := []float64{0.0007708155495898285, 0.0001682518420433508, 0.0003258109387130754}
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(Phi_want[i], got.Phi[i]) {
			t.Errorf("Phi / got %v expected %v", got.Phi, Phi_want)
			break
		}
	}
}
