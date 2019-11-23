package api

import (
	"fmt"
	"math"
)

type Phase int

const (
	Single Phase = 1
	Two    Phase = 2
)

type XYV struct {
	X_ []float64
	Y_ []float64
	V  float64
}

type FlashResult struct {
	X_   []float64 `json:"x"`
	Y_   []float64 `json:"y"`
	V    float64   `json:"v"`
	Vliq float64   `json:"Vliq"`
	Vvap float64   `json:"Vvap"`
}

func Flash_Init(components Comps, P float64, T float64, z_ []float64) (ph Phase, fraction XYV, err error) {
	bp := BublP_init(components, T, z_)
	bpRes, err := BublP(components, Eq_Input{T, bp.P, z_, bp.Y})
	if err != nil {
		return 0, XYV{}, err
	}

	err = nil
	dp := DewP_init(components, T, z_)
	dpRes, err := DewP(components, Eq_Input{T, dp.P, dp.X_, z_})
	if err != nil {
		return 0, XYV{}, err
	}
	fmt.Printf("P: %v\nPd: %v\nPb: %v\n", P, dpRes.P, bpRes.P)
	if P < dpRes.P || P > bpRes.P {
		ph = Single
		return ph, XYV{}, nil
	}

	ph = Two
	fraction.V = (bpRes.P - P) / (bpRes.P - dpRes.P)
	nc := len(components.Data)
	fraction.X_ = make([]float64, nc)
	fraction.Y_ = make([]float64, nc)
	for i := 0; i < nc; i++ {
		fraction.X_[i] = (1-fraction.V)*z_[i] + fraction.V*dpRes.X_[i]
		fraction.Y_[i] = (1-fraction.V)*bpRes.Y_[i] + fraction.V*z_[i]
	}
	return ph, fraction, nil
}

func Flash(components Comps, P float64, T float64, z_ []float64) (res FlashResult, err error) {
	ph, xyv, err := Flash_Init(components, P, T, z_)
	if err != nil {
		return FlashResult{}, err
	}
	if ph != Two {
		return FlashResult{}, nil
	}
	maxit := 3000
	nc := len(z_)
	K := make([]float64, nc)

	for i := 0; i < maxit; i++ {
		gvi_L := GetVolumeInput{P, T, xyv.X_, "L"}
		Vliq, err := GetVolume(components, gvi_L)
		if err != nil {
			return FlashResult{}, err
		}
		phi_L, fug_L, err := Fugacity(components, NewtonInput{Vliq, P, T, xyv.X_})
		if err != nil {
			return FlashResult{}, err
		}
		gvi_V := GetVolumeInput{P, T, xyv.Y_, "V"}
		Vvap, err := GetVolume(components, gvi_V)
		if err != nil {
			return FlashResult{}, err
		}
		phi_V, fug_V, err := Fugacity(components, NewtonInput{Vvap, P, T, xyv.Y_})
		if err != nil {
			return FlashResult{}, err
		}
		for j := 0; j < nc; j++ {
			K[i] = phi_L[i] / phi_V[i]
		}
		vold := xyv.V
		for iv := 0; iv < 100; iv++ {
			F := 0.
			dFdv := 0.
			for k := 0; k < nc; k++ {
				F += z_[k] * (K[k] - 1) / (1 + xyv.V*(K[k]-1))
				dFdv += -z_[k] * Pow(K[k]-1, 2) / Pow((1+xyv.V*(K[k]-1)), 2)
			}
			delv := -F / dFdv
			if math.Abs(delv) < 1e-5 {
				break
			}
			xyv.V += delv
		}

		// check converged
		converged := true
		for j := 0; j < nc; j++ {
			if math.Abs(fug_L[j]-fug_V[j]) > 1e-5 {
				converged = false
			}
		}
		if math.Abs(xyv.V-vold) < 1e-5 && converged {
			res.Vliq = Vliq
			res.Vvap = Vvap
			break
		}

		// update x,y
		for j := 0; j < nc; j++ {
			xyv.X_[j] = z_[i] / (1 + xyv.V*K[j] - 1)
			xyv.Y_[j] = K[j] * xyv.X_[j]
		}
	}
	res.X_ = xyv.X_
	res.Y_ = xyv.Y_
	res.V = xyv.V
	return
}
