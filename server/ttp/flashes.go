package ttp

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"time"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/jdb"
	"github.com/leesungbin/PC-SAFT/server/ternary"
)

type chanFlashErr struct {
	data FlashResult
	err  bool
}
type chanFlashInput struct {
	Z_ []float64
}

func Flashes_ttp(jDB jdb.DB, w http.ResponseWriter, r *http.Request) {
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
		case <-time.After(400 * time.Millisecond):
			// jsonDatas = append(jsonDatas, Eq_Result{})
		}
	}
	var names []string
	for _, d := range comps.Data {
		names = append(names, d.Name)
	}
	// jsonDatas sort
	x := func(p1, p2 *FlashResult) bool {
		x1, _, _ := ternary.Abc2xy(p1.X_[0], p1.X_[1], p1.X_[2])
		x2, _, _ := ternary.Abc2xy(p2.X_[0], p2.X_[1], p2.X_[2])
		return x1 < x2
	}

	x_rev := func(p1, p2 *FlashResult) bool {
		x1, _, _ := ternary.Abc2xy(p1.X_[0], p1.X_[1], p1.X_[2])
		x2, _, _ := ternary.Abc2xy(p2.X_[0], p2.X_[1], p2.X_[2])
		return x1 > x2
	}

	By(x).Sort(jsonDatas)

	// initial value
	x0, y0, _ := ternary.Abc2xy(jsonDatas[0].X_[0], jsonDatas[0].X_[1], jsonDatas[0].X_[2])
	rev_flag := false

	for i := 0; i < len(jsonDatas)-1; i++ {
		x1, y1, _ := ternary.Abc2xy(jsonDatas[i].X_[0], jsonDatas[i].X_[1], jsonDatas[i].X_[2])
		x2, y2, _ := ternary.Abc2xy(jsonDatas[i+1].X_[0], jsonDatas[i+1].X_[1], jsonDatas[i+1].X_[2])
		min := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
		for j := i + 2; j < len(jsonDatas); j++ {
			x2, y2, _ := ternary.Abc2xy(jsonDatas[j].X_[0], jsonDatas[j].X_[1], jsonDatas[j].X_[2])
			tmp := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
			check_rev := math.Sqrt((x2-x0)*(x2-x0) + (y2-y0)*(y2-y0))
			if tmp > check_rev {
				rev_flag = true
				break
			}
			if tmp < min {
				piece := jsonDatas[j]
				copy(jsonDatas[i+1:j+1], jsonDatas[i:j])
				jsonDatas[i] = piece
			}
		}
	}

	if rev_flag {
		By(x_rev).Sort(jsonDatas)

		for i := 0; i < len(jsonDatas)-1; i++ {
			x1, y1, _ := ternary.Abc2xy(jsonDatas[i].X_[0], jsonDatas[i].X_[1], jsonDatas[i].X_[2])
			x2, y2, _ := ternary.Abc2xy(jsonDatas[i+1].X_[0], jsonDatas[i+1].X_[1], jsonDatas[i+1].X_[2])
			min := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
			for j := i + 2; j < len(jsonDatas); j++ {
				x2, y2, _ := ternary.Abc2xy(jsonDatas[j].X_[0], jsonDatas[j].X_[1], jsonDatas[j].X_[2])
				tmp := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))

				if tmp < min {
					piece := jsonDatas[j]
					copy(jsonDatas[i+1:j+1], jsonDatas[i:j])
					jsonDatas[i] = piece
				}
			}
		}
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

type By func(p1, p2 *FlashResult) bool

func (by By) Sort(flashResults []FlashResult) {
	ps := &flashResultsSorter{
		flashResults: flashResults,
		by:           by,
	}
	sort.Sort(ps)
}

type flashResultsSorter struct {
	flashResults []FlashResult
	by           func(p1, p2 *FlashResult) bool
}

func (s *flashResultsSorter) Len() int {
	return len(s.flashResults)
}

func (s *flashResultsSorter) Swap(i, j int) {
	s.flashResults[i], s.flashResults[j] = s.flashResults[j], s.flashResults[i]
}

func (s *flashResultsSorter) Less(i, j int) bool {
	return s.by(&s.flashResults[i], &s.flashResults[j])
}
