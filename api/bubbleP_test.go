package api

import (
	"testing"
)

// at T = 338.7 K
// Component 1  =  ethane
//  Mw (g/mol)  Tc (K)   Pc (bar)   omega     Tb (K)
//    30.070    305.40     48.800   0.099     184.60
//      m       sig (A)  epsk (K)     kAB      eAB (K)  dipole (D)  xp
//    1.6069    3.5206    191.42    0.000000     0.00    0.000    0.00000

// Component 2  =  n-hexane
//  Mw (g/mol)  Tc (K)   Pc (bar)   omega     Tb (K)
//    86.178    507.50     30.100   0.299     341.90
//      m       sig (A)  epsk (K)     kAB      eAB (K)  dipole (D)  xp
//    3.0576    3.7983    236.77    0.000000     0.00    0.000    0.00000

// Component 3  =  cyclohexane
//  Mw (g/mol)  Tc (K)   Pc (bar)   omega     Tb (K)
//    84.162    553.50     40.700   0.212     353.80
//      m       sig (A)  epsk (K)     kAB      eAB (K)  dipole (D)  xp
//    2.5303    3.8499    278.11    0.000000     0.00    0.000    0.00000
// @@@BubbleP_ini start :  338.7 [0, 0.2, 0.3, 0.5]
// @@@BubbleP_init end :  18.065865443773628 [0, 0.9669884021380635, 0.015248381227162324, 0.017763216634774145]

var want_BublP_init = Result{
	18.065865443773628,
	[]float64{0.9669884021380635, 0.015248381227162324, 0.017763216634774145},
}
var want_BublP = CalculationResult{}

var Composition_NNN = []float64{0.2, 0.3, 0.5}
var Temperature = 338.7

func Test_BublP_init(t *testing.T) {
	got := NNN_ethane_nHexane_cyclohexane.BublP_init(Composition_NNN, Temperature)
	if !PassWithAccuracy4(got.P, want_BublP_init.P) {
		t.Errorf("Expected %v got %v", want_BublP_init.P, got.P)
	}
	for i, v := range got.y {
		if !PassWithAccuracy4(v, want_BublP_init.y[i]) {
			t.Errorf("Expected %v got %v", want_BublP_init.y[i], v)
		}
	}
}

func Test_BublP(t *testing.T) {
	got := NNN_ethane_nHexane_cyclohexane.BublP(CalculationInput{T: Temperature, x_: Composition_NNN})
	t.Errorf("%v", got)
}
