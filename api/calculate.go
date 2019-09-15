package api

import (
	"errors"
	"fmt"
	"math"
)

type Component struct {
	Name string
	// Data order : mw, Tc, Pc, omega, Tb, m, sig, eps, k, e, d, x
	// idx        : 0   1   2   3      4   5  6    7    8  9  10 11
	mw    float64
	Tc    float64
	Pc    float64
	omega float64
	Tb    float64
	m     float64
	sig   float64
	eps   float64
	k     float64
	e     float64
	d     float64
	x     float64
}
type Comps struct {
	data []Component
	phi  []float64
	Z    float64
}

type CrossAssociatedValues struct {
	eAB [][]float64
	kAB [][]float64
}

type FindVolumeInput struct {
	P  float64
	T  float64
	z_ []float64 // x_ or y_
}

func PrepareCrossParameter(components Comps) (res CrossAssociatedValues) {
	// # of components <- expected to 3
	nc := len(components.data)
	kAB := make([][]float64, nc)
	eAB := make([][]float64, nc)

	for i := 0; i < nc; i++ {
		kAB[i] = make([]float64, nc)
		eAB[i] = make([]float64, nc)
	}

	// fill cross association by Walbach-Sandler combining rules
	for i := 0; i < nc; i++ {
		kAB[i][i] = components.data[i].k
		eAB[i][i] = components.data[i].e
		// fmt.Println(kAB[i][i], eAB[i][i])
	}

	var sig_i, sig_j float64
	for i := 0; i < nc; i++ {
		sig_i = ConvertAtoM_sig(&components.data[i].sig)
		for j := i + 1; j < nc; j++ {
			sig_j = ConvertAtoM_sig(&components.data[j].sig)
			kAB[i][j] = math.Sqrt(kAB[i][i]*kAB[j][j]) * math.Pow(math.Sqrt(sig_i*sig_j)*2/(sig_i+sig_j), 3)
			kAB[j][i] = kAB[i][j]
			eAB[i][j] = (eAB[i][i] + eAB[j][j]) / 2
			eAB[j][i] = eAB[i][j]
			// fmt.Printf("kAB[i][j] : %f, eAB[i][j] : %f\n", kAB[i][j], eAB[i][j])
		}
	}
	res = CrossAssociatedValues{eAB: eAB, kAB: kAB}
	return res
}

// sig, convert (A) to (m)
func ConvertAtoM_sig(sig *float64) float64 {
	if *sig > 1e-5 {
		*sig = *sig * 1e-10
		return *sig
	} else {
		return *sig
	}
}

// math.Pow(x,y) x**y
func Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

type NewtonInput struct {
	V0 float64
	P  float64
	T  float64
	z_ []float64
}

func (components *Comps) Peos_P(in NewtonInput) (f float64) {
	res := components.PCsaft(PCsaftInput{in.V0, in.T, in.z_})
	Peos := R * in.T / in.V0 * res.Z
	f = Peos - in.P
	return f
}

func (components *Comps) FindV_newton(in NewtonInput) (Vres float64, err error) {
	max_iter := 100
	V := in.V0
	dV := V * 1e-5

	for i := 0; i < max_iter; i++ {
		f := components.Peos_P(in)
		if math.Abs(f/in.P) < 1e-5 {
			return V, nil
		}
		in.V0 = V + dV
		dfdV := (components.Peos_P(in) - f) / dV
		if math.Abs(dfdV*V/in.P) < 1e-5 {
			return -1, errors.New("Convergence error")
		}
		delV := -f / dfdV
		V += delV
	}
	return V, nil
}

func (components *Comps) GetVolume(in FindVolumeInput) (V float64) {
	// initial guess with ideal gas equation
	V0 := R * in.T / in.P
	V, err := components.FindV_newton(NewtonInput{V0, in.P, in.T, in.z_})
	if err != nil {
		panic(err)
	}
	return V
}

func (components *Comps) Fugacity(V float64, P float64, z_ []float64) (fug []float64) {
	if len(components.phi) == 1 && (components.phi)[0] == 0 {
		fmt.Println("0 fugacity")
	}
	for i := 0; i < len(z_); i++ {
		fug[i] = (components.phi)[i] * z_[i] * P
	}
	return fug
}
