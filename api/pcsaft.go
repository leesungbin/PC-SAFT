package api

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

type PCsaftInput struct {
	V  float64   `json:"V"`
	T  float64   `json:"T"`
	X_ []float64 `json:"x"`
}

// returns fugacity coefficient & compressibility factor
type PCsaftResult struct {
	Phi []float64 `json:"phi"`
	Z   float64   `json:"Z"`
}

const N_av float64 = 6.02214179e23
const Pi float64 = math.Pi
const R float64 = 8.314e-5 // (m3.bar/mol/K)

// PCsaft equation of state
func PCsaft(components Comps, C PCsaftInput) (res PCsaftResult, err error) {
	// start := time.Now()
	// fmt.Printf("PCsaft input : %v\n", C)
	nc := len(components.Data) // # of components
	rho_m := 1 / C.V           // molar density
	rho_num := rho_m * N_av    // number density
	M := 0.                    // mean chain length
	d := make([]float64, nc)   // effective segmemt diameter (m)
	zet := make([]float64, 4)
	g := make([]float64, nc)
	rho_dg := make([]float64, nc)
	ek_AB := PrepareCrossParameter(components)

	for i := 0; i < nc; i++ {
		// _ = ConvertAtoM_sig(&components.data[i].sig)
		M += C.X_[i] * components.Data[i].M
		d[i] = components.Data[i].Sig * (1 - 0.12*math.Exp(-3*components.Data[i].Eps/C.T))

		for j := 0; j < 4; j++ {
			zet[j] += Pi / 6 * rho_num * C.X_[i] * components.Data[i].M * Pow(d[i], float64(j))
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

		rho_dg[i] = zet[3]/Pow(1-zet[3], 2) +
			d[i]/2*(3*zet[2]/Pow(1-zet[3], 2)+6*zet[2]*zet[3]/Pow(1-zet[3], 3)) +
			+Pow(d[i]/2, 2)*(4*Pow(zet[2], 2)/Pow(1-zet[3], 3)+6*zet[2]*zet[2]*zet[3]/Pow(1-zet[3], 4))

		Zhc += -C.X_[i] * (components.Data[i].M - 1) / g[i] * rho_dg[i]
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

	// Binary parameters,, assume 0
	components.Keps = make([][]float64, nc)
	for i := 0; i < nc; i++ {
		components.Keps[i] = make([]float64, nc)
	}

	for i := 0; i < nc; i++ {
		for j := 0; j < nc; j++ {
			sig_ij3 := Pow((components.Data[i].Sig+components.Data[j].Sig)/2., 3)
			e_kT := math.Sqrt(components.Data[i].Eps*components.Data[j].Eps) * (1 - components.Keps[i][j]) / C.T
			m2esig3 += C.X_[i] * C.X_[j] * components.Data[i].M * components.Data[j].M * e_kT * sig_ij3
			m2e2sig3 += C.X_[i] * C.X_[j] * components.Data[i].M * components.Data[j].M * e_kT * e_kT * sig_ij3
		}
	}
	Zdisp := -2*Pi*rho_num*dI1*m2esig3 - Pi*rho_num*M*(C1*dI2+C2*n*I2)*m2e2sig3

	// ## association term [not completed]
	Zassoc := 0.
	ncas := 0
	for i := 0; i < nc; i++ {
		if components.Data[i].K > 1e-10 {
			ncas++
			break
		}
	}

	var idx int8                 // to catch only associating component's idx
	var F1, K1, Del1, X1 float64 // for ncas == 1
	F := make([][]float64, nc)   // belows are for ncas >=2
	K := make([][]float64, nc)
	DEL := make([][]float64, nc)
	roDDEL := make([][]float64, nc)
	X := make([]float64, nc)

	if ncas == 1 {
		for i := 0; i < nc; i++ {
			if ek_AB.E_AB[i][i] > 1e-10 {
				idx = int8(i)
			}
		}
		// fmt.Printf("%v %v\n", ek_AB.eAB, idx)
		F1 = math.Exp(ek_AB.E_AB[idx][idx]/C.T) - 1
		K1 = Pow(components.Data[idx].Sig, 3) * ek_AB.K_AB[idx][idx]
		Del1 = g[idx] * F1 * K1
		// fmt.Printf("%v %v %v\n", F1, K1, Del1)
		if C.X_[idx] < 1e-10 { // to avoid math error for pure substance
			X1 = 1.
			Zassoc = 0
		} else {
			X1 = (-1 + math.Sqrt(1+4*rho_num*C.X_[idx]*Del1)) / (2 * rho_num * C.X_[idx] * Del1)
			roDDel := rho_dg[idx] * F1 * K1
			DX := -C.X_[idx] * Pow(X1, 3) / (1 + rho_num*C.X_[idx]*X1*X1*Del1) * (Del1 + roDDel)
			Zassoc = C.X_[idx] * 2 * (1/X1 - 1./2.) * rho_num * DX
		}
		// ## for fugaticy calculation

	} else if ncas >= 2 {
		for i := 0; i < nc; i++ {
			F[i] = make([]float64, nc)
			K[i] = make([]float64, nc)
			DEL[i] = make([]float64, nc)
			roDDEL[i] = make([]float64, nc)
		}

		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				F[i][j] = math.Exp(ek_AB.E_AB[i][j]/C.T) - 1
				K[i][j] = Pow((components.Data[i].Sig+components.Data[j].Sig)/2, 3) * ek_AB.K_AB[i][j]
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
		for i := range X {
			X[i] = 1
		}

		// # guess of X for associating component
		for i := 0; i < nc; i++ {
			if ek_AB.E_AB[i][i] > 1e-10 && C.X_[i] > 1e-10 {
				X[i] = (-1 + math.Sqrt(1+4*rho_num*C.X_[i]*DEL[i][i])) / (2 * rho_num * C.X_[i] * DEL[i][i])
			}
		}
		for {
			Xold := X
			err := 0.
			for i := 0; i < nc; i++ {
				sum := 1.
				for j := 0; j < nc; j++ {
					sum += rho_num * C.X_[i] * X[j] * DEL[i][j]
				}
				X[i] = 1 / sum
				err += math.Abs(X[i] - Xold[i])
			}
			if err/float64(nc+1) < 1e-5 {
				break
			}
		}

		// iteration
		// DX := make([]float64, nc)
		// for {
		// 	DXold := DX
		// 	err := 0.
		// 	for i := 0; i < nc; i++ {
		// 		sum := 0.
		// 		for j := 0; j < nc; j++ {
		// 			sum += C.x_[j] * X[j] * (DEL[i][j] + roDDEL[i][j])
		// 			if j == i {
		// 				continue
		// 			}
		// 			sum += rho_num * C.x_[j] * DX[j] * DEL[i][j]
		// 		}
		// 		DX[i] = -X[i] * X[i] / (1 + rho_num*C.x_[i]*X[i]*X[i]*DEL[i][i]) * sum
		// 		err += math.Abs(DX[i] - DXold[i])
		// 	}
		// 	if err/float64(nc+1) < 1e-6 {
		// 		break
		// 	}
		// }

		// matrix
		A := make([][]float64, nc)
		B := make([]float64, nc)
		for i := 0; i < nc; i++ {
			A[i] = make([]float64, nc)
		}
		for i := 0; i < nc; i++ {
			A[i][i] = 1
			B[i] = 0
			for j := 0; j < nc; j++ {
				A[i][j] += X[i] * X[i] * rho_num * C.X_[j] * DEL[i][j]
				B[i] += -X[i] * X[i] * C.X_[j] * X[j] * (DEL[i][j] + roDDEL[i][j])
			}
		}
		a := make([]float64, nc*nc)
		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				a[i*nc+j] = A[i][j]
			}
		}

		A_ := mat.NewDense(nc, nc, a)
		B_ := mat.NewDense(nc, 1, B)
		var DX mat.Dense
		DX.Solve(A_, B_)

		Zassoc = 0
		for i := 0; i < nc; i++ {
			Zassoc += C.X_[i] * 2 * (1/X[i] - 1./2.) * rho_num * DX.At(i, 0) // matrix
			// Zassoc += C.x_[i] * 2 * (1/X[i] - 1./2.) * rho_num * DX[i]
		}
	}

	// ## polar term - Jog and Chapman [Complete]

	Zpolar := 0.
	nc_pol := 0 // # number of polar components
	mu2k := make([]float64, nc)

	// belows are used when calculating fugacity
	var A2p, A3p, I2p, I3p, DI2p, DI3p, sumij, sumijl float64
	dd := make([][]float64, nc)

	for i := 0; i < nc; i++ {
		if components.Data[i].D > 1e-10 {
			nc_pol += 1
			mu2k[i] = Pow(components.Data[i].D, 2) / 1.380572e26 // # dipol^2 (D^2) / k    (m3 K)
		}
	}
	if nc_pol > 0 {
		for i := range dd {
			dd[i] = make([]float64, nc)
			for j := 0; j < nc; j++ {
				dd[i][j] = (d[i] + d[j]) / 2
			}
		}

		rosta := 6 / Pi * zet[3]
		I2p = (1 - 0.3618*rosta - 0.3205*Pow(rosta, 2) + 0.1078*Pow(rosta, 3)) / Pow(1-0.5236*rosta, 2)
		I3p = (1 + 0.62378*rosta - 0.11658*Pow(rosta, 2)) / (1 - 0.59056*rosta + 0.20059*Pow(rosta, 2))
		DI2p = (0.6854 - 0.83043848*rosta + 0.3234*Pow(rosta, 2) - 0.05644408*Pow(rosta, 3)) / Pow(1-0.5236*rosta, 3)
		DI3p = (1.21434 - 0.63434*rosta - 0.0562765454*Pow(rosta, 2)) / Pow(1-0.59056*rosta+0.20059*Pow(rosta, 2), 2)
		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				sumij += C.X_[i] * C.X_[j] * components.Data[i].M * components.Data[j].M *
					components.Data[i].X * components.Data[j].X * mu2k[i] * mu2k[j] / Pow(dd[i][j], 3)
			}
		}
		A2p = -2 * Pi / 9 * rho_num / Pow(C.T, 2) * sumij * I2p
		roDA2p := A2p - 2*Pi/9*rho_num/Pow(C.T, 2)*sumij*rosta*DI2p
		for i := 0; i < nc; i++ {
			for j := 0; j < nc; j++ {
				for l := 0; l < nc; l++ {
					sumijl += C.X_[i] * C.X_[j] * C.X_[l] * components.Data[i].M * components.Data[j].M * components.Data[l].M * components.Data[i].X * components.Data[j].X * components.Data[l].X * mu2k[i] * mu2k[j] * mu2k[l] / (dd[i][j] * dd[j][l] * dd[i][l])
				}
			}
		}
		A3p = 5 * Pi * Pi / 162 * Pow(rho_num, 2) / Pow(C.T, 3) * sumijl * I3p
		roDA3p := 2*A3p + 5*Pi*Pi/162*Pow(rho_num, 2)/Pow(C.T, 3)*sumijl*rosta*DI3p

		if math.Abs(A2p) < 1e-10 {
			Zpolar = 0
		} else {
			Zpolar = 2/(1-A3p/A2p)*roDA2p - 1/Pow(1-A3p/A2p, 2)*(roDA2p-roDA3p)
		}
	}
	Z := 1. + Zhc + Zdisp + Zassoc + Zpolar // calculated
	if Z < 0 {
		return res, errors.New(fmt.Sprintf("Z < 0\nZhc : %v, Zdisp : %v, Zassoc : %v, Zpolar : %v\n", Zhc, Zdisp, Zassoc, Zpolar))
	}
	res.Z = Z // calculated compressibility factor (Z)

	// ## calculate fugacity ------------------------------------------------------------------------------------------
	Ahs := 1 / zet[0] * (3*zet[1]*zet[2]/(1-zet[3]) + Pow(zet[2], 3)/zet[3]/Pow(1-zet[3], 2) +
		(Pow(zet[2], 3)/Pow(zet[3], 2)-zet[0])*math.Log(1-zet[3]))
	Ahc := M * Ahs
	for i := 0; i < nc; i++ {
		Ahc += -C.X_[i] * (components.Data[i].M - 1) * math.Log(g[i])
	}
	tx := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		tx[i] = make([]float64, nc)
		for j := 0; j < nc; j++ {
			tx[i][j] = Pi / 6 * rho_num * components.Data[j].M * Pow(d[j], float64(i))
		}
	}
	Ahsx := make([]float64, nc)
	for k := 0; k < nc; k++ {
		Ahsx[k] = -tx[0][k]/zet[0]*Ahs + 1/zet[0]*(3*(tx[1][k]*zet[2]+zet[1]*tx[2][k])/(1-zet[3])+
			3*zet[1]*zet[2]*tx[3][k]/Pow(1-zet[3], 2)+3*Pow(zet[2], 2)*tx[2][k]/zet[3]/Pow(1-zet[3], 2)+
			Pow(zet[2], 3)*tx[3][k]*(3*zet[3]-1)/Pow(zet[3], 2)/Pow(1-zet[3], 3)+
			((3*Pow(zet[2], 2)*tx[2][k]*zet[3]-2*Pow(zet[2], 3)*tx[3][k])/Pow(zet[3], 3)-tx[0][k])*
				math.Log(1-zet[3])+(zet[0]-Pow(zet[2], 3)/Pow(zet[3], 2))*tx[3][k]/(1-zet[3]))
	}
	// fmt.Printf("%v %v %v\n", Ahc, tx, Ahsx) // complete
	gx := make([][]float64, nc)
	for i := 0; i < nc; i++ {
		gx[i] = make([]float64, nc)
		for k := 0; k < nc; k++ {
			gx[i][k] = tx[3][k]/Pow(1-zet[3], 2) +
				d[i]/2*(3*tx[2][k]/Pow(1-zet[3], 2)+6*zet[2]*tx[3][k]/Pow(1-zet[3], 3)) +
				Pow(d[i]/2, 2)*(4*zet[2]*tx[2][k]/Pow(1-zet[3], 3)+6*zet[2]*zet[2]*tx[3][k]/Pow(1-zet[3], 4))
		}
	}
	// fmt.Printf("%v %v\n", Ahsx, gx) // complete
	Ahcx := make([]float64, nc)
	for k := 0; k < nc; k++ {
		Ahcx[k] = components.Data[k].M*Ahs + M*Ahsx[k] - (components.Data[k].M-1)*math.Log(g[k]) // # with correlation
		for i := 0; i < nc; i++ {
			Ahcx[k] += -C.X_[i] * (components.Data[i].M - 1) / g[i] * gx[i][k]
		}
	}

	lnphi := make([]float64, nc)
	for k := 0; k < nc; k++ {
		lnphi[k] = res.Z - 1 - math.Log(res.Z) + Ahc + Ahcx[k]
		for j := 0; j < nc; j++ {
			lnphi[k] += -C.X_[j] * Ahcx[j]
		}
	}
	// fmt.Printf("%v %v %v\n", gx, Ahcx, lnphi) // complete

	// # Dispersion term
	Adisp := -2*Pi*rho_num*I1*m2esig3 - Pi*rho_num*M*C1*I2*m2e2sig3

	m2ekTsig3x := make([]float64, nc)
	m2ekT2sig3x := make([]float64, nc)
	for k := 0; k < nc; k++ {
		for j := 0; j < nc; j++ {
			sig3 := Pow((components.Data[k].Sig+components.Data[j].Sig)/2, 3)
			ekT := math.Sqrt(components.Data[k].Eps*components.Data[j].Eps) * (1 - components.Keps[k][j]) / C.T
			m2ekTsig3x[k] += 2 * components.Data[k].M * C.X_[j] * components.Data[j].M * ekT * sig3
			m2ekT2sig3x[k] += 2 * components.Data[k].M * C.X_[j] * components.Data[j].M * Pow(ekT, 2) * sig3
		}
	}
	// fmt.Printf("%v %v %v\n", Adisp, m2ekTsig3x, m2ekT2sig3x) // complete

	C1x := make([]float64, nc)
	for k := 0; k < nc; k++ {
		C1x[k] = C2*tx[3][k] - Pow(C1, 2)*(components.Data[k].M*(8*n-2*Pow(n, 2))/Pow(1-n, 4)-
			components.Data[k].M*(20*n-27*Pow(n, 2)+12*Pow(n, 3)-2*Pow(n, 4))/Pow(1-n, 2)/Pow(2-n, 2))
	}

	ax := make([][]float64, 7)
	bx := make([][]float64, 7)
	for i := 0; i < 7; i++ {
		ax[i] = make([]float64, nc)
		bx[i] = make([]float64, nc)
		for k := 0; k < nc; k++ {
			ax[i][k] = components.Data[k].M/(M*M)*a1[i] + components.Data[k].M/(M*M)*(3-4/M)*a2[i]
			bx[i][k] = components.Data[k].M/(M*M)*b1[i] + components.Data[k].M/(M*M)*(3-4/M)*b2[i]
		}
	}
	// fmt.Printf("%v %v %v\n", C1x, ax, bx) // complete
	I1x := make([]float64, nc)
	I2x := make([]float64, nc)
	for k := 0; k < nc; k++ {
		for i := 0; i < 7; i++ {
			I1x[k] += a[i]*float64(i)*tx[3][k]*Pow(n, float64(i-1)) + ax[i][k]*Pow(n, float64(i))
			I2x[k] += b[i]*float64(i)*tx[3][k]*Pow(n, float64(i-1)) + bx[i][k]*Pow(n, float64(i))
		}
	}
	Adispx := make([]float64, nc)
	for k := 0; k < nc; k++ {
		Adispx[k] = -2*Pi*rho_num*(I1x[k]*m2esig3+I1*m2ekTsig3x[k]) -
			Pi*rho_num*((components.Data[k].M*C1*I2+M*C1x[k]*I2+M*C1*I2x[k])*m2e2sig3+
				M*C1*I2*m2ekT2sig3x[k])
	}
	// fmt.Printf("%v %v %v\n", I1x, I2x, Adispx) // complete
	for k := 0; k < nc; k++ {
		lnphi[k] += Adisp + Adispx[k]
		for j := 0; j < nc; j++ {
			lnphi[k] += -C.X_[j] * Adispx[j]
		}
	}
	// fmt.Printf("%v\n", lnphi) // complete

	// # Association term
	// # single associaing components of I
	if ncas == 1 {
		Aassoc := C.X_[idx] * 2 * (math.Log(X1) - X1/2. + 1./2.)
		Delx := make([]float64, nc)
		Xx := make([]float64, nc)
		for k := 0; k < nc; k++ {
			Delx[k] = gx[idx][k] * F1 * K1
			Xx[k] = -Pow(X1, 2) / (1 + rho_num*C.X_[idx]*Pow(X1, 2)*Del1) * (rho_num * C.X_[idx] * X1 * Delx[k])
		}
		Xx[idx] += -Pow(X1, 2) / (1 + rho_num*C.X_[idx]*Pow(X1, 2)*Del1) * (rho_num * X1 * Del1)
		// fmt.Printf("%v\n", Xx) // complete
		Aassocx := make([]float64, nc)
		for k := 0; k < nc; k++ {
			Aassocx[k] = C.X_[idx] * 2 * (1./X1 - 1./2.) * Xx[k]
		}
		Aassocx[idx] += 2 * (math.Log(X1) - X1/2. + 1./2.)
		// fmt.Printf("%v\n", Aassocx) // complete
		// fmt.Printf("%v\n", lnphi) // complete
		// fmt.Printf("%v %v\n", Aassoc, Aassocx)
		for k := 0; k < nc; k++ {
			lnphi[k] += Aassoc + Aassocx[k]
			for j := 0; j < nc; j++ {
				lnphi[k] += -C.X_[j] * Aassocx[j]
			}
		}
		// fmt.Printf("%v\n", lnphi)
	} else if ncas >= 2 {
		// DELx = [ [ [0]*ncP for j in range(ncP) ] for i in range(ncP) ]
		DELx := make([][][]float64, nc)
		for i := 0; i < nc; i++ {
			DELx[i] = make([][]float64, nc)
			for j := 0; j < nc; j++ {
				DELx[i][j] = make([]float64, nc)
				for k := 0; k < nc; k++ {
					Gijxk := tx[3][k]/Pow(1-zet[3], 2) +
						d[i]*d[j]/(d[i]+d[j])*(3*tx[2][k]/Pow(1-zet[3], 2)+6*zet[2]*tx[3][k]/Pow(1-zet[3], 3)) +
						Pow(d[i]*d[j]/(d[i]+d[j]), 2)*(4*zet[2]*tx[2][k]/Pow(1-zet[3], 3)+6*Pow(zet[2], 2)*tx[3][k]/Pow(1-zet[3], 4))
					DELx[i][j][k] = Gijxk * F[i][j] * K[i][j]
				}
			}
		}
		Xx := make([][]float64, nc)
		for i := 0; i < nc; i++ {
			Xx[i] = make([]float64, nc)
		}
		for {
			Xxold := Xx
			err := 0.
			for i := 0; i < nc; i++ {
				for k := 0; k < nc; k++ {
					sum := rho_num * X[k] * DEL[i][k]
					for j := 0; j < nc; j++ {
						sum += rho_num * C.X_[j] * X[j] * DELx[i][j][k]
						if j == i {
							continue
						}
						sum += rho_num * C.X_[j] * Xx[j][k] * DEL[i][j]
					}

					Xx[i][k] = -Pow(X[i], 2) / (1 + rho_num*C.X_[i]*Pow(X[i], 2)*DEL[i][i]) * sum
					err += math.Abs(Xx[i][k] - Xxold[i][k])
				}
			}
			if err/float64(nc+1)/float64(nc+1) < 1e-6 {
				break
			}
		}
		Aassoc := 0.
		for i := 0; i < nc; i++ {
			Aassoc += C.X_[i] * 2 * (math.Log(X[i]) - X[i]/2 + 1/2)
		}
		Aassocx := make([]float64, nc)
		for k := 0; k < nc; k++ {
			for i := 0; i < nc; i++ {
				Aassocx[k] += C.X_[i] * 2 * (1/X[i] - 1/2) * Xx[i][k]
			}
			Aassocx[k] += 2 * (math.Log(X[k]) - X[k]/2 + 1/2)
		}
		// # add to fugacity coefficient
		for k := 0; k < nc; k++ {
			lnphi[k] += Aassoc + Aassocx[k]
			for j := 0; j < nc; j++ {
				lnphi[k] += -C.X_[j] * Aassocx[j]
			}
		}

	}
	if nc_pol > 0 && math.Abs(A2p) > 1e-10 {
		Apolar := A2p / (1 - A3p/A2p)
		dI2px := make([]float64, nc)
		dI3px := make([]float64, nc)
		A2px := make([]float64, nc)
		A3px := make([]float64, nc)
		Apolx := make([]float64, nc)
		for k := 0; k < nc; k++ {
			dI2px[k] = DI2p * rho_num * components.Data[k].M * Pow(d[k], 3)
			dI3px[k] = DI3p * rho_num * components.Data[k].M * Pow(d[k], 3)
		}
		for k := 0; k < nc; k++ {
			sum := 0.
			for i := 0; i < nc; i++ {
				sum += C.X_[i] * components.Data[i].M * components.Data[k].M * components.Data[i].X * components.Data[k].X * mu2k[i] * mu2k[k] / Pow(dd[i][k], 3)
			}
			A2px[k] = -2 * Pi / 9 * rho_num / Pow(C.T, 2) * (2*sum*I2p + sumij*dI2px[k])
		}
		for k := 0; k < nc; k++ {
			sum := 0.
			for i := 0; i < nc; i++ {
				for j := 0; j < nc; j++ {
					sum += C.X_[i] * C.X_[j] * components.Data[i].M * components.Data[j].M * components.Data[k].M *
						components.Data[i].X * components.Data[j].X * components.Data[k].X * mu2k[i] * mu2k[j] * mu2k[k] / (dd[i][j] * dd[j][k] * dd[i][k])
				}
			}
			A3px[k] = 5 * Pow(Pi, 2) / 162 * Pow(rho_num, 2) / Pow(C.T, 3) * (3*sum*I3p + sumijl*dI3px[k])
		}
		for k := 0; k < nc; k++ {
			Apolx[k] = 2/(1-A3p/A2p)*A2px[k] - 1/Pow(1-A3p/A2p, 2)*(A2px[k]-A3px[k])
		}
		for k := 0; k < nc; k++ {
			lnphi[k] += Apolar + Apolx[k]
			for j := 0; j < nc; j++ {
				lnphi[k] += -C.X_[j] * Apolx[j]
			}
		}
	}
	phi := make([]float64, nc)
	for k := 0; k < nc; k++ {
		phi[k] = math.Exp(lnphi[k])
	}
	res.Phi = phi
	// components.phi = phi
	// components.Z = res.Z

	// elapsed := time.Since(start)

	// fmt.Println("--- PC-SAFT calculation results ---")
	// fmt.Printf("time : %v\n", elapsed)
	// fmt.Printf("phi : %v\n", res.Phi)
	// fmt.Printf("PCsaft res : %v\nZhc : %v, Zdisp : %v, Zassoc : %v, Zpolar : %v\n", Z, Zhc, Zdisp, Zassoc, Zpolar)
	// fmt.Println("-----------------------------------")
	return
}
