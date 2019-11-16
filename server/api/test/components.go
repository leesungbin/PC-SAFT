package test

import (
	. "github.com/leesungbin/PC-SAFT/server/api"
)

var Ethane = Component{Name: "ethane", Mw: 30.070, Tc: 305.40, Pc: 48.800, Omega: 0.099, Tb: 184.60, M: 1.6069, Sig: 3.5206, Eps: 191.42, K: 0.000000, E: 0.00, D: 0.000, X: 0.00000}
var Nhexane = Component{Name: "n-hexane", Mw: 86.178, Tc: 507.50, Pc: 30.100, Omega: 0.299, Tb: 341.90, M: 3.0576, Sig: 3.7983, Eps: 236.77, K: 0.000000, E: 0.00, D: 0.000, X: 0.00000}
var Cyclohexane = Component{Name: "cyclohexane", Mw: 84.162, Tc: 553.50, Pc: 40.700, Omega: 0.212, Tb: 353.80, M: 2.5303, Sig: 3.8499, Eps: 278.11, K: 0.000000, E: 0.00, D: 0.000, X: 0.00000}
var Ethanol_polar = Component{Name: "ethanol (polar)", Mw: 46.069, Tc: 513.9, Pc: 61.4, Omega: 0.644, Tb: 351.4, M: 2.2049, Sig: 3.2774, Eps: 187.24, K: 0.03363, E: 2652.7, D: 1.7, X: 0.29466}
var Water_polar = Component{Name: "water (polar)", Mw: 18.015, Tc: 647.3, Pc: 221.2, Omega: 0.344, Tb: 373.15, M: 1.0405, Sig: 2.9657, Eps: 175.15, K: 0.08924, E: 2706.7, D: 1.85, X: 0.66245}
var Benzene = Component{Name: "benzene", Mw: 78.114, Tc: 562.2, Pc: 48.9, Omega: 0.212, Tb: 353.2, M: 2.4653, Sig: 3.6478, Eps: 287.35, K: 0.0, E: 0.0, D: 0.0, X: 0.0}
var Methanol_polar = Component{Name: "methanol (polar)", Mw: 32.042, Tc: 512.6, Pc: 80.9, Omega: 0.556, Tb: 337.7, M: 1.7266, Sig: 3.1369, Eps: 168.84, K: 0.06311, E: 2585.9, D: 1.7, X: 0.35128}
var AceticAcid = Component{Name: "acetic acid", Mw: 60.052, Tc: 592.7, Pc: 57.9, Omega: 0.447, Tb: 391.1, M: 1.3403, Sig: 3.8582, Eps: 211.59, K: 0.07555, E: 3044.4, D: 0, X: 0}

var NNN_ethane_nHexane_cyclohexane = Comps{Data: []Component{Ethane, Nhexane, Cyclohexane}}
var NNP_ethane_nHexane_ethanol = Comps{Data: []Component{Ethane, Nhexane, Ethanol_polar}}
var PPP_methanol_water_aceticacid = Comps{Data: []Component{Methanol_polar, Water_polar, AceticAcid}}

var PPP_Pressure = 1.013                 // kPa
var PPP_Temperarute = 345.07893999999993 // K
var PPP_composition = []float64{0.7919, 0.208, 0.0001}
