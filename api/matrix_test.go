package api

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func Test_matrix(t *testing.T) {
	data := make([]float64, 36)
	for i := 0; i < len(data); i++ {
		data[i] = float64(i)
	}
	// | 4 ,   1 | | x |     | 4 |
	// |         | |   |  =  |   |
	// | 3 ,   1 | | y |     | 3 |
	a := []float64{4, 1, 3, 1}
	A := mat.NewDense(2, 2, a)
	b := []float64{4, 3}
	B := mat.NewDense(2, 1, b)
	var c mat.Dense
	c.Solve(A, B)
	t.Errorf("%v\n", c)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
