package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	. "github.com/leesungbin/PC-SAFT/api"
	"github.com/leesungbin/PC-SAFT/env"
	"github.com/leesungbin/PC-SAFT/parser"

	"strings"

	_ "github.com/lib/pq"
)

type Service struct {
	db *sql.DB
}

func main() {
	var s *Service
	env := env.GetAppEnv()
	// values := fmt.Sprintln(env)
	// fmt.Println(values)
	// prop := fmt.Sprintf("dbname=%s sslmode=disable", env.POSTGRES_DBNAME)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", env.POSTGRES_URL, env.USER, env.PASS, env.POSTGRES_DBNAME)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("db connection failed")
		panic(err)
	} else {
		fmt.Println("db connected.")
		s = &Service{db: db}
	}

	port := fmt.Sprintf(":%s", env.PORT)
	fmt.Printf("PORT%s\n", port)
	log.Fatal(http.ListenAndServe(port, s))

}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	url := strings.Split(r.URL.Path, "/")
	// fmt.Printf("%v\n", url)
	switch url[1] {

	default:
		fmt.Fprintf(w, "PC-SAFT api server")
		return
	// case "init":
	// 	success := schema.AddPreparedDB(db)
	// 	if success {
	// 		fmt.Fprintf(w, "Prepared DB added successfully.")
	// 	} else {
	// 		fmt.Fprintf(w, "Failed to add prepared DB.")
	// 	}
	// 	return
	case "version":
		res_json := map[string]string{"version": "1.0"}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return

	case "bublp":
		if r.Method == "POST" {
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

			res, err := comps.BublP(BP_Input{T: res_parse.T, X_: res_parse.X_})
			if err != nil {
				res_json := map[string]interface{}{"status": 0, "error": err}
				print, _ := json.Marshal(res_json)
				fmt.Fprintf(w, "%s", print)
				return
			}

			data, _ := json.Marshal(res)
			if err != nil {
				res_json := map[string]string{"status": "marshal failed"}
				print, _ := json.Marshal(res_json)
				fmt.Fprintf(w, "%s", print)
				return
			}
			res_json := map[string]interface{}{"data": fmt.Sprintf("\"%s\"", data)}
			print, _ := json.Marshal(res_json)
			fmt.Fprintf(w, "%s", print)
		} else {
			res_json := map[string]string{"msg": "GET request is not supprted"}
			print, _ := json.Marshal(res_json)
			fmt.Fprintf(w, "%s", print)
		}
		return
	case "equil":
		if r.Method == "POST" {

		}
		return
	}
}
