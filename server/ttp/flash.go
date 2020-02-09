package ttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/jdb"
)

func Flash_ttp(jDB jdb.DB, w http.ResponseWriter, r *http.Request) {
	var j jsonInput
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	res, err := Flash(comps, j.P, j.T, j.X_)
	if err != nil {
		res_json := map[string]interface{}{"status": 0, "error": err}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return
	}

	res_json := map[string]FlashResult{"data": res}
	print, _ := json.Marshal(res_json)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", print)

	return
}
