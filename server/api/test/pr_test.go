package test

import (
	"testing"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_PR_ab(t *testing.T) {
	amix, bmix := PR_ab(NNN_ethane_nHexane_cyclohexane, Temperature, Composition_NNN)
	got := []float64{amix, bmix}
	want := []float64{2.5752384454956447e-05, 8.479628858481546e-05}
	if !PassWithAccuracy4(got[0], want[0]) && !PassWithAccuracy4(got[1], want[1]) {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

// passed
func Test_PPP_PR_ab(t *testing.T) {
	amix, bmix := PR_ab(PPP_methanol_water_aceticacid, PPP_Temperarute, PPP_composition)
	got := []float64{amix, bmix}
	want := []float64{1.3614718649951733e-05, 3.639932381244384e-05}
	if !PassWithAccuracy4(got[0], want[0]) && !PassWithAccuracy4(got[1], want[1]) {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func Test_PR_vol(t *testing.T) {
	Vvap, Vliq := PR_vol(NNN_ethane_nHexane_cyclohexane, Pressure, Temperature, Composition_NNN)
	// just for liq volume..!
	got := []float64{Vvap, Vliq}
	want := 0.00011024625418372402 // liq volume
	if !PassWithAccuracy4(got[1], want) {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

// passed for liquid
func Test_PPP_PR_vol(t *testing.T) {
	Vvap, Vliq := PR_vol(PPP_methanol_water_aceticacid, PPP_Pressure, PPP_Temperarute, PPP_composition)
	// just for liq volume..!
	got := []float64{Vvap, Vliq}
	want := 4.4658595745981564e-05 // liq volume
	if !PassWithAccuracy4(got[1], want) {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}
