package jdb

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"
	"strings"

	"github.com/leesungbin/PC-SAFT/server/api"
)

type DB map[string][](api.RowForm)

func (dbs DB) FilterWithName(n string) (res [](api.RowForm)) {
	rdb := dbs["datas"]
	for _, db := range rdb {
		if strings.Contains(n, db.Data.Name) {
			res = append(res, db)
		}
	}
	return res
}
func New() *DB {
	return new(DB)
}
func (dbs *DB) Read() error {
	// jsonFile, err := os.Open("data.json")
	// defer jsonFile.Close()
	// if err != nil {
	// 	return fmt.Errorf("file open error %v", err)
	// }
	// jsondata, err := ioutil.ReadAll(jsonFile)
	// if err != nil {
	// 	return fmt.Errorf("file read error: %v", err)
	// }

	err := json.Unmarshal([]byte(JSON), &dbs)
	if err != nil {
		return fmt.Errorf("unmarshal error : %v", err)
	}
	return nil
}
