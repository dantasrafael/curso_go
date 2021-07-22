package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres dbname=curso_golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(`INSERT INTO usuarios(id, nome) VALUES($1, $2)`)

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
