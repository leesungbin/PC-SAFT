package ttp

import (
	"errors"
	"fmt"
	"strings"
)

func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

type info struct {
	Nc    int
	query string
	mode  int // Eq: 1, BP: 2, BT: 3, DP: 4, DT: 5,
}

func getInfoFromBody(j jsonInput) (res info, err error) {
	if len(j.Id) == 0 {
		err = errors.New("No components.")
		return
	}
	err = nil
	if len(j.X_) == 0 && len(j.Y_) == 0 {
		res.mode = 1
	} else if len(j.X_) != 0 {
		if j.T != 0 {
			res.mode = 2
		} else {
			res.mode = 3
		}
	} else {
		if j.T != 0 {
			res.mode = 4
		} else {
			res.mode = 5
		}
	}
	res.Nc = len(j.Id)
	res.query = fmt.Sprintf("select * from component where id in (%s)", arrayToString(j.Id, ","))
	return
}
