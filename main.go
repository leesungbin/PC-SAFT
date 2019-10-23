package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

	. "github.com/leesungbin/PC-SAFT/api"
	"github.com/leesungbin/PC-SAFT/env"
	"github.com/leesungbin/PC-SAFT/schema"

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
	db, err := sql.Open("postgres", "dbname=pcsaft sslmode=disable")

	if err != nil {
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

	case "init":
		success := schema.AddPreparedDB(db)
		if success {
			fmt.Fprintf(w, "Prepared DB added successfully.")
		} else {
			fmt.Fprintf(w, "Failed to add prepared DB.")
		}
		return

	case "version":
		fmt.Fprintf(w, "{\"version\":\"1.0\"}")
		return

	case "bublp":
		if r.Method == "POST" {
			r.ParseForm()
			form := r.Form

			raw_id := fmt.Sprintf("%v", form["id"][0])
			nc := len(strings.Split(raw_id, ","))

			T, _ := strconv.ParseFloat(form["T"][0], 64)
			z := strings.Split(form["x"][0], ",")
			x := make([]float64, nc)

			for i := 0; i < len(x); i++ {
				x[i], _ = strconv.ParseFloat(z[i], 64)
			}

			query := fmt.Sprintf("select * from component where id in (%v);", raw_id)
			rows, err := db.Query(query)
			if err != nil {
				fmt.Printf("%v\n", err)
			}

			var comps Comps
			comps.Data = make([]Component, nc)
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
				// fmt.Printf("component : %v\n", component)
				comps.Data[i] = component
			}
			res, err := comps.BublP(BP_Input{T: T, X_: x})
			// fmt.Printf("%v\n", comps)
			if err != nil {
				fmt.Fprintf(w, "{\n\"status\": 0}\n")
			}
			data, _ := json.Marshal(res)
			if err != nil {
				fmt.Fprintf(w, "{\n\"status\": \"marshal failed\"}\n")
			}
			fmt.Fprintf(w, "{\n\"data\" : %s\n}", data)

		} else {
			fmt.Fprintf(w, "{\"msg\":\"GET request is not supported\"}")
		}
		return
	}
}
