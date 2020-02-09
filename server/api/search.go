package api

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type RowForm struct {
	Id   int       `json:"id"`
	Data Component `json:"data"`
}

func SearchWithName(db *sql.DB, name string) (comps []RowForm, err error) {
	i := strings.Index(name, ";")
	if i > -1 {
		err = errors.New("sql injection")
		return
	}
	query := fmt.Sprintf("select * from component where name like '%%%s%%'", name)
	rows, err_query := db.Query(query)
	if err_query != nil {
		err = err_query
		return
	}

	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var row RowForm
		if err_scan := rows.Scan(
			&row.Id, &row.Data.Name, &row.Data.Mw, &row.Data.Tc, &row.Data.Pc, &row.Data.Omega,
			&row.Data.Tb, &row.Data.M, &row.Data.Sig, &row.Data.Eps, &row.Data.K,
			&row.Data.E, &row.Data.D, &row.Data.X); err_scan != nil {
			err = err_scan
			return
		}
		comps = append(comps, row)
	}
	return
}

// SearchWithName 에서 name만 없을 뿐.. 함수를 통합할 고민 중
func SearchAll(db *sql.DB) (comps []RowForm, err error) {
	query := fmt.Sprintf("select * from component")
	rows, err_query := db.Query(query)
	if err_query != nil {
		err = err_query
		return
	}

	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var row RowForm
		if err_scan := rows.Scan(
			&row.Id, &row.Data.Name, &row.Data.Mw, &row.Data.Tc, &row.Data.Pc, &row.Data.Omega,
			&row.Data.Tb, &row.Data.M, &row.Data.Sig, &row.Data.Eps, &row.Data.K,
			&row.Data.E, &row.Data.D, &row.Data.X); err_scan != nil {
			err = err_scan
			return
		}
		comps = append(comps, row)
	}
	return
}

// db로부터의 메소드 기능 사용하지 않을 것
// SearchWithName []RowForm, err
// SearchAll []RowForm, err
