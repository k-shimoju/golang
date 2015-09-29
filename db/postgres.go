// postgres
package main

import (
	"./model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres password=pass dbname=test sslmode=disable")
	//sql := "select * from test"
	sql := "select * from test where id = $1 AND value = $2"
	rows, err := db.Query(sql, 1, "test")
	defer db.Close()
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var r model.Record

		err := rows.Scan(&r.Id, &r.Value)
		if err != nil {
			panic(err)
		}
		fmt.Println(strconv.Itoa(r.Id) + " " + r.Value)
	}
}
