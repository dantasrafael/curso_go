package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	connStr := "user=postgres password=postgres dbname=curso_golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios(
		id INTEGER GENERATED BY DEFAULT AS IDENTITY,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)

}