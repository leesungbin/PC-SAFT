package test

import (
	"testing"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_Flash(t *testing.T) {
	// 300.        0.143        0.75 0.2 0.05
	res, err := Flash(PPP_methanol_water_aceticacid, 0.143, 300, []float64{0.75, 0.2, 0.05})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Errorf("%v\n", res)
}

func Test_Flash_Init(t *testing.T) {
	ph, xyv, err := Flash_Init(PPP_methanol_water_aceticacid, 0.143, 300, []float64{0.75, 0.2, 0.05})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Errorf("%v %v\n", ph, xyv)
}
