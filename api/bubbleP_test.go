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

var Ethane = Component{"ethane", 30.070, 305.40, 48.800, 0.099, 184.60, 1.6069, 3.5206, 191.42, 0.000000, 0.00, 0.000, 0.00000}
var Nhexane = Component{"n-hexane", 86.178, 507.50, 30.100, 0.299, 341.90, 3.0576, 3.7983, 236.77, 0.000000, 0.00, 0.000, 0.00000}
var Cyclohexane = Component{"cyclohexane", 84.162, 553.50, 40.700, 0.212, 353.80, 2.5303, 3.8499, 278.11, 0.000000, 0.00, 0.000, 0.00000}
var Ethanol_polar = Component{"ethanol (polar)", 46.069, 513.9, 61.4, 0.644, 351.4, 2.2049, 3.2774, 187.24, 0.03363, 2652.7, 1.7, 0.29466}
var Water_polar = Component{"water (polar)", 18.015, 647.3, 221.2, 0.344, 373.15, 1.0405, 2.9657, 175.15, 0.08924, 2706.7, 1.85, 0.66245}
var Benzene = Component{"benzene", 78.114, 562.2, 48.9, 0.212, 353.2, 2.4653, 3.6478, 287.35, 0.0, 0.0, 0.0, 0.0}

var bublP_init_expectOutput = Result{
	18.065865443773628,
	[]float64{0.9669884021380635, 0.015248381227162324, 0.017763216634774145},
}
var bublP_expectedOutput = CalculationResult{}
var TestComponents = &Comps{data: []Component{Ethane, Nhexane, Cyclohexane}}
var PolarTestComponents = &Comps{data: []Component{Ethane, Nhexane, Ethanol_polar}}

var testComposition = []float64{0.2, 0.3, 0.5}
var TestTemperature = 338.7

func Test_BublP_init(t *testing.T) {
	got := TestComponents.BublP_init(testComposition, TestTemperature)
	if !PassWithAccuracy4(got.P, bublP_init_expectOutput.P) {
		t.Errorf("Expected %v got %v", bublP_init_expectOutput.P, got.P)
	}
	for i, v := range got.y {
		if !PassWithAccuracy4(v, bublP_init_expectOutput.y[i]) {
			t.Errorf("Expected %v got %v", bublP_init_expectOutput.y[i], v)
		}
	}
}

func Test_BublP(t *testing.T) {
	got := TestComponents.BublP(CalculationInput{T: TestTemperature, x_: testComposition})
	t.Errorf("%v", got)
}
