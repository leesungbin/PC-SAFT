package schema

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Component struct {
	Name string
	// Data order : mw, Tc, Pc, omega, Tb, m, sig, eps, k, e, d, x
	Data []float64
}

func readData(filename string) Component {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	name, _, err := reader.ReadLine()
	reader.ReadLine()
	thermoData, _, err := reader.ReadLine()
	reader.ReadLine()
	saftData, _, err := reader.ReadLine()

	data := make([]float64, 12)
	cnt := 0
	for i, d := range strings.Fields(string(thermoData)) {
		data[i], _ = strconv.ParseFloat(d, 64)
		cnt++
	}

	for j, d := range strings.Fields(string(saftData)) {
		data[cnt+j], _ = strconv.ParseFloat(d, 64)
		// if j == 1 && data[cnt+j] > 1e-5 { // : sig, convert (A) to (m) -> db error just zero value
		// 	data[cnt+j] *= 1e-10
		// }
	}

	return Component{strings.TrimSpace(string(name)), data}
}

func listOfFiles(path string) []Component {
	var res []Component
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		filepath := fmt.Sprintf("%s%s", path, f.Name())
		component := readData(filepath)
		res = append(res, component)
	}
	return res
}

func AddPreparedDB(db *sql.DB) bool {
	var dbLength int16
	err := db.QueryRow("SELECT count(*) from component").Scan(&dbLength)
	if err != nil {
		panic(err)
	}

	if dbLength < 112 { // default data length, if some data added manually don't delete data
		fmt.Printf("%d data exist.\n", dbLength)
		if dbLength > 0 {
			_, err := db.Exec("DELETE from component;") // delete all data
			if err != nil {
				panic(err)
			} else {
				fmt.Println("Successfully deleted.")
			}
		}
		components := listOfFiles("../components")
		query := "INSERT INTO component VALUES "
		for i, comp := range components {
			// fmt.Printf("%s %v\n", comp.Name, comp.Data)
			query = fmt.Sprintf("%s (%d, '%s', %s)", query, i, comp.Name, arrayWithComma(comp.Data))
			if i+1 != len(components) {
				query = fmt.Sprintf("%s,", query)
			}
		}
		query = fmt.Sprintf("%s;", query)

		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Added Prepared Data.")
			return true
		}
	} else {
		// some data added manually. protect this data
		return false
	}
}
func arrayWithComma(datas []float64) string {
	var str string
	for i, d := range datas {
		str = fmt.Sprintf("%s%f", str, d)
		if i+1 != len(datas) {
			str = fmt.Sprintf("%s,", str)
		}
	}
	return str
}
