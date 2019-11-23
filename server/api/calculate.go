package api

import (
	"errors"
	"fmt"
	"math"
)

type CrossAssociatedValues struct {
	E_AB [][]float64 `json:"eAB"`
	K_AB [][]float64 `json:"kAB"`
}

type GetVolumeInput struct {
	P     float64   `json:"P"`
	T     float64   `json:"T"`
	Z_    []float64 `json:"z"`     // x_ or y_
	State string    `json:"state"` // V or L
}

func PrepareCrossParameter(components Comps) (res CrossAssociatedValues) {
	// # of components <- expected to 3
	nc := len(components.Data)
	kAB := make([][]float64, nc)
	eAB := make([][]float64, nc)

	for i := 0; i < nc; i++ {
		kAB[i] = make([]float64, nc)
		eAB[i] = make([]float64, nc)
	}

	// fill cross association by Walbach-Sandler combining rules
	for i := 0; i < nc; i++ {
		kAB[i][i] = components.Data[i].K
		eAB[i][i] = components.Data[i].E
	}

	var sig_i, sig_j float64
	for i := 0; i < nc; i++ {
		sig_i = ConvertAtoM_sig(&components.Data[i].Sig)
		for j := i + 1; j < nc; j++ {
			sig_j = ConvertAtoM_sig(&components.Data[j].Sig)
			kAB[i][j] = math.Sqrt(kAB[i][i]*kAB[j][j]) * math.Pow(math.Sqrt(sig_i*sig_j)*2/(sig_i+sig_j), 3)
			kAB[j][i] = kAB[i][j]
			eAB[i][j] = (eAB[i][i] + eAB[j][j]) / 2
			eAB[j][i] = eAB[i][j]
		}
	}
	res = CrossAssociatedValues{E_AB: eAB, K_AB: kAB}
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
	V  float64   `json:"V"`
	P  float64   `json:"P"`
	T  float64   `json:"T"`
	Z_ []float64 `json:"z"`
}

func Peos_P(components Comps, in NewtonInput) (f float64, err error) {
	// fmt.Printf("	Peos_P input : %v\n", in)
	res, err := PCsaft(components, PCsaftInput{in.V, in.T, in.Z_})
	// fmt.Printf("		PCsaft res : %v\n", res)
	if err != nil {
		err = errors.New(fmt.Sprintf("Peos_P got error while PCsaft : %v\n", err))
		return f, err
	}
	Peos := R * in.T / in.V * res.Z
	f = Peos - in.P
	// Log(fmt.Sprintf("Peos_P : %v %v", Peos, in.P))
	return f, nil
}

func FindV_newton(components Comps, in NewtonInput) (Vres float64, err error) {
	var f, f_next float64
	max_iter := 100
	V := in.V
	dV := V * 1e-5
	for i := 0; i < max_iter; i++ {
		f, err = Peos_P(components, NewtonInput{V, in.P, in.T, in.Z_})
		if err != nil {
			return V, err
		}

		if math.Abs(f/in.P) < 1e-2 { // 정확도 1e-5까지 가는 경우가 잘 없다.
			// fmt.Printf("iterated for %d times\nfindV_newton end : V, : %v\nconverged rate : %v\n", i, V, f/in.P)
			return V, nil
		}

		V = V + dV
		f_next, err = Peos_P(components, NewtonInput{V, in.P, in.T, in.Z_})
		if err != nil {
			return V, err
		}

		dfdV := (f_next - f) / dV
		if math.Abs(dfdV*V/in.P) < 1e-5 {
			return V, errors.New("Convergence error")
		}
		delV := -f / dfdV * 0.5
		V += delV
	}
	// fmt.Printf("iterated for %d times\nfindV_newton end : V, : %v\nconverged rate : %v\n", max_iter, V, f/in.P)
	return V, nil
}

func GetVolume(components Comps, in GetVolumeInput) (V float64, err error) {
	Vvap, Vliq := PR_vol(components, in.P, in.T, in.Z_)

	var V0 float64
	if in.State == "V" {
		V0 = Vvap
	} else {
		V0 = Vliq * LiquidScale // set scalVl0 = 0.95
	}
	// Log(fmt.Sprintf("GetVolume : %v", V0))
	V, err = FindV_newton(components, NewtonInput{V0, in.P, in.T, in.Z_})
	// Log(fmt.Sprintf("After Find V_newton : %v", V))
	if err != nil {
		return V, errors.New(err.Error())
	}
	return V, nil
}

func Fugacity(components Comps, in NewtonInput) (phi, fug []float64, err error) {
	res, err := PCsaft(components, PCsaftInput{in.V, in.T, in.Z_})
	// fmt.Printf("fugacity - pcsaft res : %v\n", res)
	if err != nil {
		return phi, fug, err
	}
	if len(res.Phi) == 1 && (res.Phi)[0] == 0 {
		// fmt.Println("0 fugacity")
		fug = []float64{0}
	} else {
		fug = make([]float64, len(res.Phi))
		for i := 0; i < len(in.Z_); i++ {
			fug[i] = (res.Phi)[i] * in.Z_[i] * in.P
		}
	}

	return res.Phi, fug, nil
}
