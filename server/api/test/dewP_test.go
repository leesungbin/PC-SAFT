package test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_DewP_PPP(t *testing.T) {
	start := time.Now()
	res, err := DewP(PPP_methanol_water_aceticacid, Eq_Input{T: 300, Y_: PPP_composition})
	elapsed := time.Since(start)
	fmt.Printf("time required : %v\n", elapsed) // 46~50 ms / python: 280~300 ms
	if err != nil {
		t.Errorf("%v\n", err)
	}
	P_want := 0.12785753435719896
	X_want := []float64{0.47270619697658, 0.4555716292774705, 0.07172239960033287}
	if !PassWithAccuracy4(res.P, P_want) {
		t.Errorf("Expected %v got %v\n", P_want, res.P)
	}
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(res.X_[i], X_want[i]) {
			t.Errorf("Expectd %v got %v\n", X_want, res.X_)
			break
		}
	}
}
