package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"database/sql"

	"strings"

	. "github.com/leesungbin/PC-SAFT/server/api"
	"github.com/leesungbin/PC-SAFT/server/env"
	"github.com/leesungbin/PC-SAFT/server/parser"
	"github.com/leesungbin/PC-SAFT/server/ternary"
	"github.com/leesungbin/PC-SAFT/server/ttp"

	_ "github.com/lib/pq"
)

type chanErr struct {
	data Eq_Result
	err  bool
}
type Service struct {
	db *sql.DB
}

const (
	envPublicDir = "web/dist"
	envIndexFile = "index.html"
)

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

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := filepath.Join(envPublicDir, filepath.Clean(r.URL.Path))
		if info, err := os.Stat(p); err != nil {
			http.ServeFile(w, r, filepath.Join(envPublicDir, envIndexFile))
			return
		} else if info.IsDir() {
			http.ServeFile(w, r, filepath.Join(envPublicDir, envIndexFile))
			return
		}
		http.ServeFile(w, r, p)
	}))

	mux.Handle("/api/version", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res_json := map[string]string{"version": "1.0"}
		print, _ := json.Marshal(res_json)
		fmt.Fprintf(w, "%s", print)
		return
	}))
	mux.Handle("/api/bublp", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ttp.BublP(db, w, r)
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	mux.Handle("/api/equil", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ttp.Equil()
			return
		}
		fmt.Fprintf(w, "Get req is not supported.")
		return
	}))
	log.Fatal(http.ListenAndServe(port, mux))

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

			res, err := BublP(comps, Eq_Input{T: res_parse.T, X_: res_parse.X_})
			if err != nil {
				res_json := map[string]interface{}{"status": 0, "error": err}
				print, _ := json.Marshal(res_json)
				fmt.Fprintf(w, "%s", print)
				return
			}

			res_json := map[string]Eq_Result{"data": res}
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
			now := time.Now()

			plots := ternary.Cover()
			nc := len(plots)
			inChan := make(chan Eq_Input, nc)
			equilDatas := make(chan chanErr, nc)
			var jsonDatas []Eq_Result

			// for BublP
			if res_parse.T != 0. {
				for i := 0; i < nc; i++ {
					go func(idx int) {
						in := <-inChan
						res, err := BublP(comps, Eq_Input{T: res_parse.T, X_: in.X_})
						if err != nil {
							equilDatas <- chanErr{data: Eq_Result{}, err: true}
						} else {
							equilDatas <- chanErr{data: res, err: false}
						}
					}(i)
				}
				for i := 0; i < nc; i++ {
					a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
					fractions := []float64{a, b, c}
					inChan <- Eq_Input{T: res_parse.T, X_: fractions}
				}
				close(inChan)
			} else if res_parse.P != 0. { // for BublT
				for i := 0; i < nc; i++ {
					go func(idx int) {
						in := <-inChan
						res, err := BublT(comps, Eq_Input{P: res_parse.P, X_: in.X_})
						if err != nil {
							equilDatas <- chanErr{data: Eq_Result{}, err: true}
						} else {
							equilDatas <- chanErr{data: res, err: false}
						}
					}(i)
				}
				for i := 0; i < nc; i++ {
					a, b, c, _ := ternary.Xy2abc(plots[i].X, plots[i].Y)
					fractions := []float64{a, b, c}
					inChan <- Eq_Input{P: res_parse.P, X_: fractions}
				}
				close(inChan)
			} else {
				// error input
				fmt.Fprintf(w, "hi equil!!")
			}

			for i := 0; i < nc; i++ {
				select {
				case normal := <-equilDatas:
					if !normal.err && normal.data.P != 0 && normal.data.T != 0 {
						jsonDatas = append(jsonDatas, normal.data)
					}
				case <-time.After(50 * time.Millisecond):
					// jsonDatas = append(jsonDatas, Eq_Result{})
				}
			}
			res_json := map[string][]Eq_Result{"data": jsonDatas}
			print, _ := json.Marshal(res_json)
			fmt.Fprintf(w, "%s", print)
			fmt.Printf("success for %.2f%% inputs\n", float64(len(jsonDatas))/float64(nc)*100)
			fmt.Printf("time required : %v\n", time.Since(now))
			return
		}
	}
}
