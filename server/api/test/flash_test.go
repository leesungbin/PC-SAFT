package test

import (
	"testing"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_Flash(t *testing.T) {
	start := time.Now()
	res, err := Flash(PPP_methanol_water_aceticacid, 0.143, 300, []float64{0.75, 0.2, 0.05})
	elapsed := time.Since(start)
	t.Errorf("time required : %v\n", elapsed) // 10~15 ms
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
