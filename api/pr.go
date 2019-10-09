package api

import (
	"math"
)

func (components *Comps) PR_ab(T float64, z []float64) (amix, bmix float64) {
	nc := len(components.data)
	a := make([]float64, nc)
	b := make([]float64, nc)
	for i := 0; i < nc; i++ {
		kapa := 0.37464 + 1.54226*components.data[i].omega - 0.26992*Pow(components.data[i].omega, 2)
		alpha := Pow((1 + kapa*(1-math.Sqrt(T/components.data[i].Tc))), 2)
		a[i] = 0.45724 * R * R * Pow(components.data[i].Tc, 2) / components.data[i].Pc * alpha
		b[i] = 0.07780 * R * components.data[i].Tc / components.data[i].Pc
	}
	for i := 0; i < nc; i++ {
		bmix += z[i] * b[i]
		for j := 0; j < nc; j++ {
			aij := 1.
			if i == j {
				aij = a[i]
			} else {
				if len(components.Keps) == 0 { // not yet calculated...?
					aij = math.Sqrt(a[i] * a[j])
				} else {
					aij = math.Sqrt(a[i]*a[j]) * (1 - components.Keps[i][j])
				}
			}
			amix += z[i] * z[j] * aij
		}
	}
	return
}

func (components *Comps) PR_vol(P, T float64, z []float64) (Vvap, Vliq float64) {
	a, b := components.PR_ab(T, z)
	// # using math formula for cubic equation  x^3 + a2 * x^2 + a1 * x + a0 = 0
	a2 := b - R*T/P
	a1 := -3*b*b + (a-2*b*R*T)/P
	a0 := b*b*b + (b*b*R*T-a*b)/P

	q := (3*a1 - a2*a2) / 9
	r := (9*a1*a2 - 27*a0 - 2*Pow(a2, 3)) / 54
	D := Pow(q, 3) + r*r

	V_ := make([]float64, 3)
	if D < 0 {
		th := math.Acos(r / math.Sqrt(-q*q*q))
		V_[0] = 2*math.Sqrt(-q)*math.Cos(th/3) - a2/3
		V_[1] = 2*math.Sqrt(-q)*math.Cos(th/3+2./3*Pi) - a2/3
		V_[2] = 2*math.Sqrt(-q)*math.Cos(th/3+4./3*Pi) - a2/3

		Vvap = math.Max(math.Max(V_[0], V_[1]), V_[2])
		Vliq = math.Min(math.Min(V_[0], V_[1]), V_[2])
	} else { // D > 0 , single real root
		s := r + math.Sqrt(D)
		if s > 0 {
			s = Pow(s, 1./3.)
		} else {
			s = -(Pow(-s, 1./3.))
		}

		t := r - math.Sqrt(D)
		if t > 0 {
			t = Pow(t, 1./3.)
		} else {
			t = -(Pow(-t, 1./3.))
		}

		Vvap = -a2/3 + s + t
		Vliq = Vvap
	}
	return
}
