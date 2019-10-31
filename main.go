package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"database/sql"

	. "github.com/leesungbin/PC-SAFT/api"
	"github.com/leesungbin/PC-SAFT/env"
	"github.com/leesungbin/PC-SAFT/parser"
	"github.com/leesungbin/PC-SAFT/ternary"

	"strings"

	_ "github.com/lib/pq"
)

type Service struct {
	db *sql.DB
}

func main() {
	var s *Service
	var dsn string
	env := env.GetAppEnv()

	if b, e := strconv.ParseBool(env.DEBUG); b && e == nil {
		dsn = fmt.Sprintf("dbname=%s sslmode=disable", env.POSTGRES_DBNAME)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", env.POSTGRES_URL, env.USER, env.PASS, env.POSTGRES_DBNAME)
	}
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

			res_json := map[string]BP_Result{"data": res}
			print, _ := json.Marshal(res_json)
			fmt.Fprintf(w, "%s", print)
		} else {
			res_json := map[string]string{"msg": "GET request is not supprted"}
			print, _ := json.Marshal(res_json)
			fmt.Fprintf(w, "%s", print)
		}
		return

	// only for ternary equilibrium
	case "equil":
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

			if res_parse.T != 0. {
				plots := ternary.Cover()
				nc := len(plots)
				resChan := make([]chan BP_Result, nc)
				equilData := make([]BP_Result, nc)
				for i := 0; i < len(plots); i++ {
					resChan[i] = make(chan BP_Result)
				}

				// parrallel processing
				for i, plot := range plots {
					go func(idx int) {
						a, b, c, _ := ternary.Xy2abc(plot.X, plot.Y)

						var timeout_channel chan bool
						go func(res chan BP_Result) {
							r, e := comps.BublP(BP_Input{T: res_parse.T, X_: []float64{a, b, c}})
							// when calculation error occurs,
							if e != nil {
								// return empty result
								resChan[idx] <- BP_Result{}
							} else {
								resChan[idx] <- r
								timeout_channel <- false
							}
						}(resChan[idx])

						// check timeout
						select {
						case _ = <-timeout_channel:
							return

							// when calculation timeout
						case <-time.After(4 * time.Second):
							// return empty result
							resChan[idx] <- BP_Result{}
							return
						}
					}(i)
				}

				for i := 0; i < nc; i++ {
					equilData[i] = <-resChan[i]
				}
				res_json := map[string][]BP_Result{"data": equilData}
				print, _ := json.Marshal(res_json)
				fmt.Fprintf(w, "%s", print)
			} else if res_parse.P != 0. {
				plots := ternary.Cover()
				resChan := make([]chan BT_Result, len(plots))
				for i := 0; i < len(plots); i++ {
					resChan[i] = make(chan BT_Result)
				}
			} else {
				// error input
			}

		}
		return
	}
}
