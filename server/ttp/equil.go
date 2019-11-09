package ttp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/parser"
	"github.com/leesungbin/PC-SAFT/server/ternary"
)

type chanErr struct {
	data Eq_Result
	err  bool
}

func Equil_ttp(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	res_parse, err_parse := parser.Post(form)

	if err_parse != nil {
		res_json := map[string]interface{}{"status": 400, "error": err_parse}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return
	}

	rows, err := db.Query(res_parse.SelectQuery)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var comps Comps
	comps.Data = make([]Component, res_parse.Nc)
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var (
			id        int
			component Component
		)
		if err := rows.Scan(
			&id, &component.Name, &component.Mw, &component.Tc, &component.Pc, &component.Omega,
			&component.Tb, &component.M, &component.Sig, &component.Eps, &component.K,
			&component.E, &component.D, &component.X); err != nil {
			fmt.Printf("err : %v\n", err)
		}
		comps.Data[i] = component
	}
	now := time.Now()

	plots := ternary.Cover()
	nc := len(plots)
	inChan := make(chan Eq_Input, nc)
	equilDatas := make(chan chanErr, nc)
	var jsonDatas []Eq_Result

	// for BublP
	if res_parse.T != 0. {
		for i := 0; i < nc; i++ {
			go func(idx int) {
				in := <-inChan
				res, err := BublP(comps, Eq_Input{T: res_parse.T, X_: in.X_})
				if err != nil {
					equilDatas <- chanErr{data: Eq_Result{}, err: true}
				} else {
					equilDatas <- chanErr{data: res, err: false}
				}
			}(i)
		}
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			inChan <- Eq_Input{T: res_parse.T, X_: fractions}
		}
		close(inChan)
	} else if res_parse.P != 0. { // for BublT
		for i := 0; i < nc; i++ {
			go func(idx int) {
				in := <-inChan
				res, err := BublT(comps, Eq_Input{P: res_parse.P, X_: in.X_})
				if err != nil {
					equilDatas <- chanErr{data: Eq_Result{}, err: true}
				} else {
					equilDatas <- chanErr{data: res, err: false}
				}
			}(i)
		}
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			inChan <- Eq_Input{P: res_parse.P, X_: fractions}
		}
		close(inChan)
	} else {
		// error input
		fmt.Fprintf(w, "hi equil!!")
	}

	for i := 0; i < nc; i++ {
		select {
		case normal := <-equilDatas:
			if !normal.err && normal.data.P != 0 && normal.data.T != 0 {
				jsonDatas = append(jsonDatas, normal.data)
			}
		case <-time.After(5 * time.Millisecond):
			// jsonDatas = append(jsonDatas, Eq_Result{})
		}
	}
	res_json := map[string][]Eq_Result{"data": jsonDatas}
	print, _ := json.Marshal(res_json)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", print)
	fmt.Printf("success for %.2f%% inputs\n", float64(len(jsonDatas))/float64(nc)*100)
	fmt.Printf("time required : %v\n", time.Since(now))
	return
}
