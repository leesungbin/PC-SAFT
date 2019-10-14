package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leesungbin/PC-SAFT/env"
	"github.com/leesungbin/PC-SAFT/schema"

	"database/sql"

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
	log.Fatal(http.ListenAndServe(port, s))

}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	url := strings.Split(r.URL.Path, "/")
	fmt.Printf("%v\n", url)
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
	}
}
