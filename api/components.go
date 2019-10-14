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
