package api

import (
	"fmt"
	"math"
)

type Component struct {
	Name string
	// Data order : mw, Tc, Pc, omega, Tb, m, sig, eps, k, e, d, x
	// idx        : 0   1   2   3      4   5  6    7    8  9  10 11
	Data []float64
}
type Components []Component

type CrossAssociatedValues struct {
	eAB [][]float64
	kAB [][]float64
}

func PrepareCrossParameter(components Components) (res CrossAssociatedValues) {
	// # of components <- expected to 3
	nc := len(components)
	kAB := make([][]float64, nc)
	eAB := make([][]float64, nc)

	for i := 0; i < nc; i++ {
		kAB[i] = make([]float64, nc)
		eAB[i] = make([]float64, nc)
	}

	// fill cross association by Walbach-Sandler combining rules
	for i := 0; i < nc; i++ {
		kAB[i][i] = components[i].Data[8]
		eAB[i][i] = components[i].Data[9]
		fmt.Println(kAB[i][i], eAB[i][i])
	}

	var sig_i, sig_j float64
	for i := 0; i < nc; i++ {
		sig_i = convertAtoM_sig(components[i].Data[6])
		for j := i + 1; j < nc; j++ {
			sig_j = convertAtoM_sig(components[j].Data[6])
			kAB[i][j] = math.Sqrt(kAB[i][i]*kAB[j][j]) * math.Pow(math.Sqrt(sig_i*sig_j)*2/(sig_i+sig_j), 3)
			kAB[j][i] = kAB[i][j]
			eAB[i][j] = (eAB[i][i] + eAB[j][j]) / 2
			eAB[j][i] = eAB[i][j]
			fmt.Printf("kAB[i][j] : %f, eAB[i][j] : %f\n", kAB[i][j], eAB[i][j])
		}
	}
	res = CrossAssociatedValues{eAB: eAB, kAB: kAB}
	return res
}

// sig, convert (A) to (m)
func convertAtoM_sig(sig float64) float64 {
	if sig > 1e-5 {
		return sig * 1e-10
	} else {
		return sig
	}
}
