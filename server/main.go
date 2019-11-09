package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

	// . "github.com/leesungbin/PC-SAFT/server/api"
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
	mux.Handle("/", craHandler)
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
		if r.Method == "POST" {
			ttp.BublP_ttp(db, w, r)
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	mux.Handle("/api/equil", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ttp.Equil_ttp(db, w, r)
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	log.Fatal(http.ListenAndServe(port, mux))
}
