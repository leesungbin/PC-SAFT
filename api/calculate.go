package api

import (
	"errors"
	"fmt"
	"math"
)

type CrossAssociatedValues struct {
	eAB [][]float64
	kAB [][]float64
}

type FindVolumeInput struct {
	P     float64
	T     float64
	z_    []float64 // x_ or y_
	state string    // V or L
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
	V  float64
	P  float64
	T  float64
	z_ []float64
}

func (components *Comps) Peos_P(in NewtonInput) (f float64) {
	// fmt.Printf("	Peos_P input : %v\n", in)
	res, err := components.PCsaft(PCsaftInput{in.V, in.T, in.z_})
	// fmt.Printf("		PCsaft res : %v\n", res)
	if err != nil {
		fmt.Printf("Peos_P got error while PCsaft : %v\n", err)
	}
	Peos := R * in.T / in.V * res.Z
	f = Peos - in.P
	// Log(fmt.Sprintf("Peos_P : %v %v", Peos, in.P))
	return f
}

func (components *Comps) FindV_newton(in NewtonInput) (Vres float64, err error) {
	var f float64
	max_iter := 100
	V := in.V
	dV := V * 1e-5
	// fmt.Printf("FindV_newton : V, dV : %v %v \n", V, dV)
	for i := 0; i < max_iter; i++ {
		// fmt.Printf("Peos-P : %v\n	converged rate : %v\n", f, f/in.P)
		f = components.Peos_P(NewtonInput{V, in.P, in.T, in.z_})
		if math.Abs(f/in.P) < 1e-5 {
			fmt.Printf("iterated for %d times\nfindV_newton end : V, : %v\nconverged rate : %v\n", i, V, f/in.P)
			return V, nil
		}
		V = V + dV
		f_next := components.Peos_P(NewtonInput{V, in.P, in.T, in.z_})
		dfdV := (f_next - f) / dV
		if math.Abs(dfdV*V/in.P) < 1e-5 {
			return V, errors.New("Convergence error")
		}
		delV := -f / dfdV * 0.95
		V += delV
		// f = f_next
	}
	fmt.Printf("iterated for %d times\nfindV_newton end : V, : %v\nconverged rate : %v\n", max_iter, V, f/in.P)
	return V, nil
}

func (components *Comps) GetVolume(in FindVolumeInput) (V float64, err error) {
	Vvap, Vliq := components.PR_vol(in.P, in.T, in.z_)

	var V0 float64
	if in.state == "V" {
		V0 = Vvap
	} else {
		V0 = Vliq * 0.99 // set scalVl0 = 0.99
	}
	// Log(fmt.Sprintf("GetVolume : %v", V0))
	V, err = components.FindV_newton(NewtonInput{V0, in.P, in.T, in.z_})
	// Log(fmt.Sprintf("After Find V_newton : %v", V))
	if err != nil {
		return V, errors.New(fmt.Sprintf("%v\n", err))
	}
	return V, nil
}

func (components *Comps) Fugacity(in NewtonInput) (phi, fug []float64) {
	res, _ := components.PCsaft(PCsaftInput{in.V, in.T, in.z_})
	if len(res.Phi) == 1 && (res.Phi)[0] == 0 {
		fmt.Println("0 fugacity")
		fug = []float64{0}
	} else {
		fug = make([]float64, len(res.Phi))
		for i := 0; i < len(in.z_); i++ {
			fug[i] = (res.Phi)[i] * in.z_[i] * in.P
		}
	}

	return res.Phi, fug
}
