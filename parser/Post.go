package parser

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type PostParsed struct {
	SelectQuery string
	Nc          int
	P           float64
	T           float64
	X_          []float64
	Y_          []float64
}

// P, T, X_, Y_ input to SelectQuery & Nc
func Post(form url.Values) (res PostParsed, err []error) {
	raw_id := fmt.Sprintf("%v", form["id"][0])
	nc := len(strings.Split(raw_id, ","))
	res.Nc = nc

	T := form["T"]
	if T != nil {
		T_tmp, err_tmp := strconv.ParseFloat(T[0], 64)
		if err != nil {
			err = append(err, errors.New(fmt.Sprintf("T strconv error : %s\n", err_tmp)))
		} else {
			res.T = T_tmp
		}
	}

	P := form["P"]
	if P != nil {
		P_tmp, err_tmp := strconv.ParseFloat(P[0], 64)
		if err != nil {
			err = append(err, errors.New(fmt.Sprintf("P strconv error : %s\n", err_tmp)))
		} else {
			res.P = P_tmp
		}
	}

	x_ := form["x"]
	if x_ != nil {
		x_tmp := strings.Split(x_[0], ",")
		if nc != len(x_tmp) {
			err = append(err, errors.New(fmt.Sprintf("you are trying to calculate %d components, but the composition length is %d.\n", nc, len(x_))))
		} else {
			var err_tmp error
			x_conv := make([]float64, nc)
			for i := 0; i < nc; i++ {
				x_conv[i], err_tmp = strconv.ParseFloat(x_tmp[i], 64)
			}
			if err_tmp != nil {
				err = append(err, errors.New(fmt.Sprintf("x composition typo error.")))
			} else {
				res.X_ = x_conv
			}
		}
	}

	y_ := form["y"]
	if y_ != nil {
		y_tmp := strings.Split(y_[0], ",")
		if nc != len(y_tmp) {
			err = append(err, errors.New(fmt.Sprintf("you are trying to calculate %d components, but the composition length is %d.\n", nc, len(x_))))
		} else {
			var err_tmp error
			y_conv := make([]float64, nc)
			for i := 0; i < nc; i++ {
				y_conv[i], err_tmp = strconv.ParseFloat(y_tmp[i], 64)
			}
			if err_tmp != nil {
				err = append(err, errors.New(fmt.Sprintf("y composition typo error.")))
			} else {
				res.Y_ = y_conv
			}
		}
	}
	res.SelectQuery = fmt.Sprintf("select * from component where id in (%v);", raw_id)
	return
}
