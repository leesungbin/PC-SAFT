package ttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/jdb"
	"github.com/leesungbin/PC-SAFT/server/ternary"
)

func Equil_ttp_sync(jDB jdb.DB, w http.ResponseWriter, r *http.Request) {
	var j jsonInput
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Printf("%v %v %v\n", j.T, j.P, j.Id)
	compInfo, err_parse := getInfoFromBody(j)

	if err_parse != nil {
		res_json := map[string]interface{}{"status": 400, "error": err_parse}
		print, _ := json.Marshal(res_json)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", print)
		return
	}

	// rows, err := db.Query(res_parse.query)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	var comps Comps
	comps.Data = make([]Component, compInfo.Nc)
	// defer rows.Close()
	// for i := 0; rows.Next(); i++ {
	// 	var (
	// 		id        int
	// 		component Component
	// 	)
	// 	if err := rows.Scan(
	// 		&id, &component.Name, &component.Mw, &component.Tc, &component.Pc, &component.Omega,
	// 		&component.Tb, &component.M, &component.Sig, &component.Eps, &component.K,
	// 		&component.E, &component.D, &component.X); err != nil {
	// 		fmt.Printf("err : %v\n", err)
	// 	}
	// 	comps.Data[i] = component
	// }
	for i, id := range compInfo.Ids {
		comps.Data[i] = jDB["data"][id].Data
	}
	now := time.Now()

	plots := ternary.Cover()
	nc := len(plots)
	var jsonDatas []Eq_Result

	min := 1.
	max := 0.
	var mode string

	// for BublP
	if j.T != 0. {
		mode = "BUBLP"
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			input := Eq_Input{T: j.T, X_: fractions}
			res, err := BublP(comps, input)
			if err != nil {
				fmt.Printf("in: %v\nout: %v\nerr: %v\n", input, res, err)
			} else {
				jsonDatas = append(jsonDatas, res)
				if min > res.P || min == 1. {
					min = res.P
				}
				if max < res.P {
					max = res.P
				}
			}
		}
	} else if j.P != 0 { // for BublT
		mode = "BUBLT"
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			input := Eq_Input{P: j.P, X_: fractions}
			res, err := BublT(comps, input)
			if err != nil {
				fmt.Printf("in: %v\nout: %v\nerr: %v\n", input, res, err)
			} else {
				jsonDatas = append(jsonDatas, res)
				if min > res.T || min == 1 {
					min = res.T
				}
				if max < res.T {
					max = res.T
				}
			}
			fmt.Printf("%v%%\n", float64(i)/float64(nc))
		}
	} else {
		mode = "error"
		// error input
		fmt.Fprintf(w, "hi equil!!")
	}

	type Range struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	}
	type resJson struct {
		Data   []Eq_Result `json:"data"`
		Header Range       `json:"header"`
		Mode   string      `json:"mode"`
		Names  []string    `json:"names"`
	}

	var names []string
	for _, d := range comps.Data {
		names = append(names, d.Name)
	}

	res_json := map[string]resJson{"result": resJson{jsonDatas, Range{min, max}, mode, names}}

	print, _ := json.Marshal(res_json)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", print)
	fmt.Printf("success for %.2f%% inputs\n", float64(len(jsonDatas))/float64(nc)*100)
	fmt.Printf("time required : %v\n", time.Since(now))
	return
}
