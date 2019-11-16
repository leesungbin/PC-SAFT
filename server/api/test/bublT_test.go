package test

import (
	"testing"
	"time"
	// "time"

	. "github.com/leesungbin/PC-SAFT/server/api"
)

func Test_BublT(t *testing.T) {
	start := time.Now()
	got, _ := BublT(PPP_methanol_water_aceticacid, Eq_Input{P: PPP_Pressure, X_: PPP_composition})
	elapsed := time.Since(start)
	t.Errorf("time required : %v\n", elapsed) // average 1.6~2 ms, python보다 10배 빠름
	// 오차 1% 미만
	if !PassWithAccuracyN(1, got.T, 340.74) {
		t.Errorf("Expected %v got %v", 340.74, got)
	}
	y := []float64{0.8992, 0.1007, 0.0000}
	for i, v := range got.Y_ {
		if !PassWithAccuracyN(1, v, y[i]) {
			t.Errorf("Expected %v got %v", y, got)
		}
	}
}
