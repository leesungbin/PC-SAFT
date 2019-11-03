package ternary

import (
	"errors"
	"math"
)

/**
(x,y) to
	 A
 /   \
B  -  C
(a,b,c)
*/
func Xy2abc(x, y float64) (a, b, c float64, err error) {
	sq3 := math.Sqrt(3)
	if y < 0 {
		err = errors.New("y < BC line\n")
	} else if y > sq3*x {
		err = errors.New("y > AB line\n")
	} else if y > -sq3*x+sq3 {
		err = errors.New("y > AC line\n")
	}
	if err == nil {
		a = 2. / sq3 * y
		b = math.Abs(x + y/sq3 - 1)
		c = math.Abs(x - y/sq3)
	}
	return
}

func Abc2xy(a, b, c float64) (x, y float64, err error) {
	sq3 := math.Sqrt(3)
	if a+b+c < 0.9999 || a+b+c > 1.0001 {
		err = errors.New("a+b+c != 1\n")
	} else if a < 0 || b < 0 || c < 0 {
		err = errors.New("some of a,b,c < 0\n")
	} else {
		y = a / 2 * sq3
		x = c + y/sq3
	}
	return
}
