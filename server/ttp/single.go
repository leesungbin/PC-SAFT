// through the process
package ttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/jdb"
)

// not api.BublP. It's for request resolving. mode 0: BublP, 1: BublT, 2: DewP, 3: DewT
func Single_ttp(jDB jdb.DB, mode int, w http.ResponseWriter, r *http.Request) {
	var j jsonInput
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	compInfo, err_parse := getInfoFromBody(j)
	// component, err_parse
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

	res := Eq_Result{}
	var err_eq error

	switch mode {
	case 0:
		res, err_eq = BublP(comps, Eq_Input{T: j.T, X_: j.X_})
	case 1:
		res, err_eq = BublT(comps, Eq_Input{P: j.P, X_: j.X_})
	case 2:
		res, err_eq = DewP(comps, Eq_Input{T: j.T, Y_: j.Y_})
	case 3:
		res, err_eq = DewT(comps, Eq_Input{P: j.P, Y_: j.Y_})
	}

	if err_eq != nil {
		res_json := map[string]interface{}{"status": 0, "error": err}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return
	}

	res_json := map[string]Eq_Result{"data": res}
	print, _ := json.Marshal(res_json)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", print)

	return
}
