package ttp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/ternary"
)

type chanFlashErr struct {
	data FlashResult
	err  bool
}
type chanFlashInput struct {
	Z_ []float64
}

func Flashes_ttp(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var j jsonInput
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Printf("%v %v %v\n", j.T, j.P, j.Id)
	res_parse, err_parse := getInfoFromBody(j)

	if err_parse != nil {
		res_json := map[string]interface{}{"status": 400, "error": err_parse}
		print, _ := json.Marshal(res_json)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", print)
		return
	}

	rows, err := db.Query(res_parse.query)
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
	inChan := make(chan chanFlashInput, nc)
	equilDatas := make(chan chanFlashErr, nc)
	var jsonDatas []FlashResult

	for i := 0; i < nc; i++ {
		go func(idx int) {
			in := <-inChan
			res, err := Flash(comps, j.P, j.T, in.Z_)
			if err != nil {
				equilDatas <- chanFlashErr{data: FlashResult{}, err: true}
			} else {
				equilDatas <- chanFlashErr{data: res, err: false}
			}
		}(i)
	}
	for i := 0; i < nc; i++ {
		a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
		fractions := []float64{a, b, c}
		inChan <- chanFlashInput{Z_: fractions}
	}
	close(inChan)

	for i := 0; i < nc; i++ {
		select {
		case normal := <-equilDatas:
			if !normal.err && normal.data.Vliq != 0 && normal.data.Vvap != 0 {
				jsonDatas = append(jsonDatas, normal.data)
			}
		case <-time.After(100 * time.Millisecond):
			// jsonDatas = append(jsonDatas, Eq_Result{})
		}
	}
	var names []string
	for _, d := range comps.Data {
		names = append(names, d.Name)
	}
	type resJson struct {
		Data  []FlashResult `json:"data"`
		Names []string      `json:"names"`
	}
	res_json := map[string]resJson{"result": resJson{jsonDatas, names}}
	print, _ := json.Marshal(res_json)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", print)
	fmt.Printf("success for %.2f%% inputs\n", float64(len(jsonDatas))/float64(nc)*100)
	fmt.Printf("time required : %v\n", time.Since(now))
	return
}
