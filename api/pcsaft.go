package api

import (
	"fmt"
	"math"
)

type PCsaftInput struct {
	V         float64
	T         float64
	x_        []float64
	y_        []float64
	component Components
}

// returns fugacity coefficient & compressibility factor
type PCsaftResult struct {
	Phi float64
	Z   float64
}

const N_av float64 = 6.02214179e23
const pi float64 = math.Pi
const R float64 = 8.314e-5 // (m3.bar/mol/K)

// PCsaft equation of state
func PCsaft(C PCsaftInput) (res PCsaftResult) {
	// start := time.Now()
	nc := len(C.component)   // # of components
	rho_m := 1 / C.V         // molar density
	rho_num := rho_m * N_av  // number density
	M := 0.                  // mean chain length
	d := make([]float64, nc) // effective segmemt diameter (m)
	zet := make([]float64, 4)
	g := make([]float64, nc)
	rho_dg := make([]float64, nc)
	ek_AB := PrepareCrossParameter(C.component)

	for i := 0; i < nc; i++ {
		// _ = ConvertAtoM_sig(&C.component[i].sig)
		M += C.x_[i] * C.component[i].m
		d[i] = C.component[i].sig * (1 - 0.12*math.Exp(-3*C.component[i].eps/C.T))

		for j := 0; j < 4; j++ {
			zet[j] += pi / 6 * rho_num * C.x_[i] * C.component[i].m * Pow(d[i], float64(j))
		}
	}

	// ## hard sphere term [Test Complete]
	// Zhc initialize with m_*Zhs
	Zhc := M * (zet[3]/(1-zet[3]) +
		3*zet[1]*zet[2]/(zet[0]*Pow(1-zet[3], 2)) +
		Pow(zet[2], 3)*(3-zet[3])/(zet[0]*Pow(1-zet[3], 3)))

	for i := 0; i < nc; i++ {
		g[i] = 1/(1-zet[3]) +
			d[i]/2*3*zet[2]/Pow(1-zet[3], 2) +
			Pow(d[i]/2, 2)*2*Pow(zet[2], 2)/Pow(1-zet[3], 3)

		// rho_dg[i] = zet[3]/Pow(1-zet[3], 2) +
		// 	d[i]/2*(3*zet[2]/Pow(1-zet[3], 2)+6*zet[2]*zet[3]/Pow(1-zet[3], 3)) +
		// 	Pow(d[i]/2, 2)*(4*Pow(zet[2], 2)/Pow(1-zet[3], 3)+6*Pow(zet[2], 2)*zet[3]/Pow(1-zet[3], 4))
		rho_dg[i] = zet[3]/Pow(1-zet[3], 2) +
			d[i]/2*(3*zet[2]/Pow(1-zet[3], 2)+6*zet[2]*zet[3]/Pow(1-zet[3], 3)) +
			+Pow(d[i]/2, 2)*(4*Pow(zet[2], 2)/Pow(1-zet[3], 3)+6*zet[2]*zet[2]*zet[3]/Pow(1-zet[3], 4))
		Zhc += -C.x_[i] * (C.component[i].m - 1) / g[i] * rho_dg[i]
	}

	// ## Dispersion term [Test Complete]
	n := zet[3]
	a0 := []float64{0.9105631445, 0.6361281449, 2.6861347891, -26.547362491, 97.759208784, -159.59154087, 91.297774084}
	a1 := []float64{-0.3084016918, 0.1860531159, -2.5030047259, 21.419793629, -65.255885330, 83.318680481, -33.746922930}
	a2 := []float64{-0.0906148351, 0.4527842806, 0.5962700728, -1.7241829131, -4.1302112531, 13.776631870, -8.6728470368}
	b0 := []float64{0.7240946941, 2.2382791861, -4.0025849485, -21.003576815, 26.855641363, 206.55133841, -355.60235612}
	b1 := []float64{-0.5755498075, 0.6995095521, 3.8925673390, -17.215471648, 192.67226447, -161.82646165, -165.20769346}
	b2 := []float64{0.0976883116, -0.2557574982, -9.1558561530, 20.642075974, -38.804430052, 93.626774077, -29.666905585}

	a := make([]float64, 7)
	b := make([]float64, 7)
	var I1, I2, dI1, dI2 float64

	for i := 0; i < 7; i++ {
		a[i] = a0[i] + (M-1)/M*a1[i] + (M-1)*(M-2)/Pow(M, 2)*a2[i]
		b[i] = b0[i] + (M-1)/M*b1[i] + (M-1)*(M-2)/Pow(M, 2)*b2[i]
		I1 += a[i] * Pow(n, float64(i))
		I2 += b[i] * Pow(n, float64(i))
		dI1 += a[i] * float64(i+1) * Pow(n, float64(i))
		dI2 += b[i] * float64(i+1) * Pow(n, float64(i))
	}
	C1 := 1 / (1 +
		M*(8*n-2*n*n)/Pow(1-n, 4) +
		(1-M)*(20*n-27*n*n+12*Pow(n, 3)-2*Pow(n, 4))/Pow((1-n)*(2-n), 2))
	C2 := -C1 * C1 *
		(M*(-4*n*n+20*n+8)/Pow(1-n, 5) +
			(1-M)*(2*Pow(n, 3)+12*n*n-48*n+40)/Pow((1-n)*(2-n), 3))
	C1 = 1 / (1 + M*(8*n-2*n*n)/Pow(1-n, 4) - (M-1)*(20*n-27*Pow(n, 2)+12*Pow(n, 3)-2*Pow(n, 4))/Pow(1-n, 2)/Pow(2-n, 2))
	C2 = -C1 * C1 * (M*(-4*Pow(n, 2)+20*n+8)/Pow(1-n, 5) - (M-1)*(2*Pow(n, 3)+12*Pow(n, 2)-48*n+40)/Pow(1-n, 3)/Pow(2-n, 3))
	m2esig3 := 0.
	m2e2sig3 := 0.

	for i := 0; i < nc; i++ {
		for j := 0; j < nc; j++ {
			sig_ij3 := Pow((C.component[i].sig+C.component[j].sig)/2, 3)
			e_kT := math.Sqrt(C.component[i].eps*C.component[j].eps) * (1 - ek_AB.kAB[i][j]) / C.T
			m2esig3 += C.x_[i] * C.x_[j] * C.component[i].m * C.component[j].m * e_kT * sig_ij3
			m2e2sig3 += C.x_[i] * C.x_[j] * C.component[i].m * C.component[j].m * Pow(e_kT, 2) * sig_ij3
		}
	}
	Zdisp := -2*pi*rho_num*dI1*m2esig3 - pi*rho_num*M*(C1*dI2+C2*n*I2)*m2e2sig3

	// ## association term
	Zassoc := 0.
	ncas := 0
	for i := 0; i < nc; i++ {
		for j := 0; j < nc; j++ {
			if ek_AB.kAB[i][j] > 1e-10 {
				ncas++
				break
			}
		}
	}

	if ncas == 1 {
		var idx int8
		for i := 0; i < nc; i++ {
			if ek_AB.eAB[i][i] > 1e-10 {
				idx = int8(i)
			}
		}
		F := math.Exp(ek_AB.eAB[idx][idx]/C.T) - 1
		K := Pow(C.component[idx].sig, 3) * ek_AB.kAB[idx][idx]
		Del := g[idx] * F * K

		if C.x_[idx] < 1e-10 { // to avoid math error for pure substance
			// X := 1.
			Zassoc = 0
		} else {
			X := (-1 + math.Sqrt(1+4*rho_num*C.x_[idx]*Del)) / (2 * rho_num * C.x_[idx] * Del)
			roDDel := rho_dg[idx] * F * K
			DX := -C.x_[idx] * Pow(X, 3) / (1 + rho_num*C.x_[idx]*Pow(X, 2)*Del) * (Del + roDDel)
			Zassoc = C.x_[idx] * 2 * (1/X - 1/2) * rho_num * DX
		}

	} else if ncas >= 2 {

		F := make([][]float64, nc)
		K := make([][]float64, nc)
		DEL := make([][]float64, nc)
		roDDEL := make([][]float64, nc)

		for i := 0; i < nc; i++ {
			F[i] = make([]float64, nc)
			K[i] = make([]float64, nc)
			DEL[i] = make([]float64, nc)
			roDDEL[i] = make([]float64, nc)
		}

		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				F[i][j] = math.Exp(ek_AB.eAB[i][j]/C.T) - 1
				K[i][j] = Pow((C.component[i].sig+C.component[j].sig)/2, 3) * ek_AB.kAB[i][j]
				g_ij := 1/(1-zet[3]) + d[i]*d[j]/(d[i]+d[j])*3*zet[2]/Pow(1-zet[3], 2) +
					Pow(d[i]*d[j]/(d[i]+d[j]), 2) + 2*Pow(zet[2], 2)/Pow(1-zet[3], 3)
				DEL[i][j] = g_ij * F[i][j] * K[i][j]
				roDg_ij := zet[3]/Pow(1-zet[3], 2) +
					d[i]*d[j]/(d[i]+d[j])*(3*zet[2]/Pow(1-zet[3], 2)+6*zet[2]*zet[3]/Pow(1-zet[3], 3)) +
					Pow(zet[i]*zet[j]/(zet[i]+zet[j]), 2)*
						(4*Pow(zet[2], 2)/Pow(1-zet[3], 3)+6*Pow(zet[2], 2)*zet[3]/Pow(1-zet[3], 4))
				roDDEL[i][j] = roDg_ij * F[i][j] * K[i][j]
			}
		}
		// # Find X's by iteration (successive substitution)
		// # defaul value X = 1 for non-associating components
		X := make([]float64, nc)
		for i := range X {
			X[i] = 1
		}

		// # guess of X for associating component
		for i := 0; i < nc; i++ {
			if ek_AB.eAB[i][i] > 1e-10 && C.x_[i] > 1e-10 {
				X[i] = (-1 + math.Sqrt(1+4*rho_num*C.x_[i]*DEL[i][i])) / (2 * rho_num * C.x_[i] * DEL[i][i])
			}
		}
		for {
			Xold := X
			err := 0.
			for i := 0; i < nc; i++ {
				sum := 1.
				for j := 0; j < nc; j++ {
					sum += rho_num * C.x_[i] * X[j] * DEL[i][j]
				}
				X[i] = 1 / sum
				err += math.Abs(X[i] - Xold[i])
			}
			if err/float64(nc+1) < 1e-5 {
				break
			}
		}
		DX := make([]float64, nc)

		for {
			DXold := DX
			err := 0.
			for i := 0; i < nc; i++ {
				sum := 0.
				for j := 0; j < nc; j++ {
					sum += C.x_[j] * X[j] * (DEL[i][j] + roDDEL[i][j])
					if j == i {
						continue
					}
					sum += rho_num * C.x_[j] * DX[j] * DEL[i][j]
				}
				DX[i] = -X[i] * X[i] / (1 + rho_num*C.x_[i]*X[i]*X[i]*DEL[i][i]) * sum
				err += math.Abs(DX[i] - DXold[i])
			}
			if err/float64(nc+1) < 1e-6 {
				break
			}
		}

		Zassoc = 0
		for i := 0; i < nc; i++ {
			Zassoc += C.x_[i] * 2 * (1/X[i] - 1/2) * rho_num * DX[i]
		}
	}

	// ## polar term - Jog and Chapman

	Zpolar := 0.
	nc_pol := 0 // # number of polar components
	mu2k := make([]float64, nc)

	for i := 0; i < nc; i++ {
		if C.component[i].d > 1e-10 {
			nc_pol += 1
			mu2k[i] = Pow(C.component[i].d, 2) / 1.380572e26 // # dipol^2 (D^2) / k    (m3 K)
		}
	}
	if nc_pol > 0 {
		dd := make([][]float64, nc)
		for i := range dd {
			dd[i] = make([]float64, nc)
			for j := 0; j < nc; j++ {
				dd[i][j] = (d[i] + d[j]) / 2
			}
		}

		rosta := 6 / pi * zet[3]
		I2p := (1 - 0.3618*rosta - 0.3205*Pow(rosta, 2) + 0.1078*Pow(rosta, 3)) / Pow(1-0.5236*rosta, 2)
		I3p := (1 + 0.62378*rosta - 0.11658*Pow(rosta, 2)) / (1 - 0.59056*rosta + 0.20059*Pow(rosta, 2))
		DI2p := (0.6854 - 0.83043848*rosta + 0.3234*Pow(rosta, 2) - 0.05644408*Pow(rosta, 3)) / Pow(1-0.5236*rosta, 3)
		DI3p := (1.21434 - 0.63434*rosta - 0.0562765454*Pow(rosta, 2)) / Pow(1-0.59056*rosta+0.20059*Pow(rosta, 2), 2)
		sumij := 0.
		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				sumij += C.x_[i] * C.x_[j] * C.component[i].m * C.component[j].m *
					C.component[i].x * C.component[j].x * mu2k[i] * mu2k[j] / Pow(dd[i][j], 3)
			}
		}
		A2p := -2 * pi / 9 * rho_num / Pow(C.T, 2) * sumij * I2p
		roDA2p := A2p - 2*pi/9*rho_num/Pow(C.T, 2)*sumij*rosta*DI2p
		sumijl := 0.
		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				for l := 0; l < nc; l++ {
					sumijl += C.x_[i] * C.x_[j] * C.x_[l] * C.component[i].m * C.component[j].m * C.component[l].m * C.component[i].x * C.component[j].x * C.component[l].x * mu2k[i] * mu2k[j] * mu2k[l] / (dd[i][j] * dd[j][l] * dd[i][l])
				}
			}
		}
		A3p := 5 * pi * pi / 162 * Pow(rho_num, 2) / Pow(C.T, 3) * sumijl * I3p
		roDA3p := 2*A3p + 5*pi*pi/162*Pow(rho_num, 2)/Pow(C.T, 3)*sumijl*rosta*DI3p

		if math.Abs(A2p) < 1e-10 {
			Zpolar = 0
		} else {
			Zpolar = 2/(1-A3p/A2p)*roDA2p - 1/Pow(1-A3p/A2p, 2)*(roDA2p-roDA3p)
		}
	}
	// ## calculate fugacity

	res.Phi = 1
	res.Z = 1. + Zhc + Zdisp + Zassoc + Zpolar
	fmt.Println("Z s : ", Zhc, Zdisp, Zassoc, Zpolar)

	// t := time.Now()
	// elapsed := t.Sub(start)
	// fmt.Printf("%v elapsed", elapsed)
	return
}
