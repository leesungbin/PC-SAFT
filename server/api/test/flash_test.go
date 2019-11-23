package test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_Flash(t *testing.T) {
	start := time.Now()
	res, err := Flash(PPP_methanol_water_aceticacid, 0.14255409405727607, 300, []float64{0.75, 0.2, 0.05})
	elapsed := time.Since(start)
	fmt.Printf("time required : %v\n", elapsed)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	X_want := []float64{0.6187836838978023, 0.30592058873122385, 0.07529572739564287}
	Y_want := []float64{0.8220163264769287, 0.14186689639453656, 0.036116777114995585}
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(res.X_[i], X_want[i]) {
			t.Errorf("x / Expected %v got %v\n", X_want[i], res.X_[i])
		}
		if !PassWithAccuracy4(res.Y_[i], Y_want[i]) {
			t.Errorf("y / Expected %v got %v\n", Y_want[i], res.Y_[i])
		}
	}
}

func Test_Flash_Init(t *testing.T) {
	ph, xyv, err := Flash_Init(PPP_methanol_water_aceticacid, 0.14255409405727607, 300, []float64{0.75, 0.2, 0.05})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if ph != Two {
		t.Errorf("not two phase")
	}
	X_want := []float64{0.61135309848829, 0.32778581463873524, 0.06086119980016644}
	Y_want := []float64{0.8192157843599228, 0.14605548336877358, 0.034716917716869744}
	V_want := 0.5
	for i := 0; i < 3; i++ {
		if !PassWithAccuracy4(xyv.X_[i], X_want[i]) {
			t.Errorf("x / Expected %v got %v\n", X_want, xyv.X_)
			break
		}
		if !PassWithAccuracy4(xyv.Y_[i], Y_want[i]) {
			t.Errorf("y / Expected %v got %v\n", Y_want, xyv.Y_)
			break
		}
	}
	if !PassWithAccuracy4(xyv.V, V_want) {
		t.Errorf("v / Expected %v got %v\n", V_want, xyv.V)
	}
}
