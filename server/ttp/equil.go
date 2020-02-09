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

type chanErr struct {
	data Eq_Result
	err  bool
}
type jsonInput struct {
	T  float64   `json:"T"`
	P  float64   `json:"P"`
	X_ []float64 `json:"x"`
	Y_ []float64 `json:"y"`
	Id []int     `json:"id"`
}

func Equil_ttp(jDB jdb.DB, w http.ResponseWriter, r *http.Request) {
	var j jsonInput
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%v %v %v\n", j.T, j.P, j.Id)
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
	inChan := make(chan Eq_Input, nc)
	equilDatas := make(chan chanErr, nc)
	var jsonDatas []Eq_Result

	min := 1.
	max := 0.
	var mode string

	// for BublP
	if j.T != 0. {
		mode = "BUBLP"
		for i := 0; i < nc; i++ {
			go func(idx int) {
				in := <-inChan
				res, err := BublP(comps, Eq_Input{T: j.T, X_: in.X_})
				if err != nil {
					equilDatas <- chanErr{data: Eq_Result{}, err: true}
				} else {
					equilDatas <- chanErr{data: res, err: false}
					if min > res.P || min == 1. {
						min = res.P
					}
					if max < res.P {
						max = res.P
					}
				}
			}(i)
		}
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			inChan <- Eq_Input{T: j.T, X_: fractions}
		}
		close(inChan)
	} else if j.P != 0 { // for BublT
		mode = "BUBLT"
		for i := 0; i < nc; i++ {
			go func(idx int) {
				in := <-inChan
				res, err := BublT(comps, Eq_Input{P: j.P, X_: in.X_})
				if err != nil {
					equilDatas <- chanErr{data: Eq_Result{}, err: true}
				} else {
					equilDatas <- chanErr{data: res, err: false}
					if (min > res.T || min == 1) && res.T > 0 {
						min = res.T
					}
					if max < res.T {
						max = res.T
					}
				}
			}(i)
		}
		for i := 0; i < nc; i++ {
			a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
			fractions := []float64{a, b, c}
			inChan <- Eq_Input{P: j.P, X_: fractions}
		}
		close(inChan)
	} else {
		mode = "error"
		// error input
		fmt.Fprintf(w, "hi equil!!")
	}

	for i := 0; i < nc; i++ {
		select {
		case normal := <-equilDatas:
			if !normal.err && normal.data.P != 0 && normal.data.T != 0 {
				jsonDatas = append(jsonDatas, normal.data)
			}
		case <-time.After(300 * time.Millisecond):
			// jsonDatas = append(jsonDatas, Eq_Result{})
		}
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
	print, err := json.Marshal(res_json)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		fmt.Fprintf(w, "{\"err\": \"marshal error : %s\"", err)
		return
	}
	fmt.Fprintf(w, "%s", print)
	fmt.Printf("success for %.2f%% inputs\n", float64(len(jsonDatas))/float64(nc)*100)
	fmt.Printf("time required : %v\n", time.Since(now))
	return
}
