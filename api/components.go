package api

type Component struct {
	Name string `json:"name"`
	// Data order : Mw, Tc, Pc, omega, Tb, m, sig, eps, k, e, d, x
	// idx        : 0   1   2   3      4   5  6    7    8  9  10 11
	Mw    float64 `json:"Mw"`
	Tc    float64 `json:"Tc"`
	Pc    float64 `json:"Pc"`
	Omega float64 `json:"w"`
	Tb    float64 `json:"Tb"`
	M     float64 `json:"m"`
	Sig   float64 `json:"sigma"`
	Eps   float64 `json:"epsilon"`
	K     float64 `json:"k"`
	E     float64 `json:"e"`
	D     float64 `json:"d"`
	X     float64 `json:"x"`
}

type Comps struct {
	Data []Component `json:"data"`
	// phi  []float64
	// Z    float64
	Keps [][]float64 `json:"Keps"`
}

var Ethane = Component{"ethane", 30.070, 305.40, 48.800, 0.099, 184.60, 1.6069, 3.5206, 191.42, 0.000000, 0.00, 0.000, 0.00000}
var Nhexane = Component{"n-hexane", 86.178, 507.50, 30.100, 0.299, 341.90, 3.0576, 3.7983, 236.77, 0.000000, 0.00, 0.000, 0.00000}
var Cyclohexane = Component{"cyclohexane", 84.162, 553.50, 40.700, 0.212, 353.80, 2.5303, 3.8499, 278.11, 0.000000, 0.00, 0.000, 0.00000}
var Ethanol_polar = Component{"ethanol (polar)", 46.069, 513.9, 61.4, 0.644, 351.4, 2.2049, 3.2774, 187.24, 0.03363, 2652.7, 1.7, 0.29466}
var Water_polar = Component{"water (polar)", 18.015, 647.3, 221.2, 0.344, 373.15, 1.0405, 2.9657, 175.15, 0.08924, 2706.7, 1.85, 0.66245}
var Benzene = Component{"benzene", 78.114, 562.2, 48.9, 0.212, 353.2, 2.4653, 3.6478, 287.35, 0.0, 0.0, 0.0, 0.0}

var NNN_ethane_nHexane_cyclohexane = &Comps{Data: []Component{Ethane, Nhexane, Cyclohexane}}
var NNP_ethane_nHexane_ethanol = &Comps{Data: []Component{Ethane, Nhexane, Ethanol_polar}}
