package ternary

import (
	"fmt"
	"testing"
)

func Test_xy2abc(t *testing.T) {
	_, _, _, err := xy2abc(2, 0)
	if err == nil {
		t.Errorf("error should not be nil\n")
	}

	a, b, c, err := xy2abc(1./2., 1./2.)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	fmt.Printf("%f %f %f\n", a, b, c)
}

func Test_abc2xy(t *testing.T) {
	x, y, err := abc2xy(0.1, 0.2, 0.7)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	fmt.Printf("%f %f\n", x, y)

}
