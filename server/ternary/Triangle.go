package ternary

import (
	"math"
)

type Point struct {
	x float64
	y float64
}
type Plot struct {
	X float64
	Y float64
	I int
}

func Cover() (data []Plot) {
	sq3 := math.Sqrt(3.)
	now := Point{0.5, sq3 / 2.}
	// 50 -> BublP : 3 sec, BublT : 12 sec ..
	n := 30.
	data = make([]Plot, int(n*(n+1)/2+1))
	dx := 1. / n
	for i := 0; ; i++ {
		if now.y < 0 {
			break
		}
		if (now.y-sq3)/(-sq3) < now.x {
			// now.x = now.y/sq3 - dx
			// now.y = sq3 * now.x
			now.y = now.y - sq3/2*dx
			now.x = now.y / sq3
			i--
			continue
		}
		data[i] = Plot{now.x, now.y, i}
		now.x += dx
	}
	return
}
