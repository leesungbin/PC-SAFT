package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/env"
	"github.com/leesungbin/PC-SAFT/server/ttp"

	_ "github.com/lib/pq"
)

// type Service struct {
// 	db *sql.DB
// }

const (
	envPublicDir = "web"
	envStaticDir = "web/static"
	envIndexFile = "index.html"
)

func main() {
	// var s *Service
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
	}

	port := fmt.Sprintf(":%s", env.PORT)
	fmt.Printf("PORT%s\n", port)

	mux := http.NewServeMux()

	craHandler := http.FileServer(http.Dir(envPublicDir))
	staticHandler := http.FileServer(http.Dir(envStaticDir))
	// bad..
	mux.Handle("/", craHandler)
	mux.Handle("/db", craHandler)
	mux.Handle("/docs", craHandler)
	mux.Handle("/static", staticHandler)

	mux.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PC-SAFT API Server\n")
		return
	}))
	mux.Handle("/api/version", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res_json := map[string]string{"version": "1.0"}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return
	}))
	mux.Handle("/api/bublp", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			ttp.BublP_ttp(db, w, r)
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	mux.Handle("/api/equil", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			ttp.Equil_ttp(db, w, r)
			// ttp.Equil_ttp_sync(db, w, r)
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	mux.Handle("/api/flash", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			ttp.Flash_ttp(db, w, r)
			return
		}
	}))
	mux.Handle("/api/datas", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			comps, err := api.SearchAll(db)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}

			res_json := map[string][](api.RowForm){"data": comps}
			print, _ := json.Marshal(res_json)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", print)
			return
		}
		fmt.Fprintf(w, "not allowed.")
		return
	}))
	mux.Handle("/api/search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type searchInput struct {
			Name string `json:"name"`
		}
		var si searchInput
		if r.Method == http.MethodPost {
			err_json := json.NewDecoder(r.Body).Decode(&si)
			if err_json != nil {
				http.Error(w, err_json.Error(), http.StatusBadRequest)
				return
			}
			comps, err := api.SearchWithName(db, si.Name)

			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}

			res_json := map[string][](api.RowForm){"data": comps}
			print, _ := json.Marshal(res_json)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", print)
			return
		}
		fmt.Fprintf(w, "not allowed.")
		return
	}))
	log.Fatal(http.ListenAndServe(port, mux))
}
