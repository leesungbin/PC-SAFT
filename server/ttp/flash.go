package ttp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/parser"
)

func Flash_ttp(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	res, err := Flash(comps, res_parse.P, res_parse.T, res_parse.X_)
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
