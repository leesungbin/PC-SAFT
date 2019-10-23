package ternary

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func Cover() {
	sq3 := math.Sqrt(3.)
	now := Point{0.5, sq3 / 2.}
	n := 100.
	dx := 1. / n
	for i := 0; ; i++ {
		if now.y < 0 {
			break
		}
		if (now.y-sq3)/(-sq3) < now.x {
			now.x = now.y/sq3 - dx
			now.y = sq3 * now.x
			continue
		}
		now.x += dx
	}
}
