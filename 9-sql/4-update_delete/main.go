package main

import (
	"database/sql"
	"fmt"
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

	// UPDATE
	stmt, err := db.Prepare(`UPDATE usuarios SET nome = $1 WHERE id = $2`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Uoxiton Istive", 1)
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := res.RowsAffected()
	fmt.Println(rows)

	// DELETE
	stmt2, err := db.Prepare(`DELETE FROM usuarios WHERE id = $1`)
	if err != nil {
		log.Fatal(err)
	}

	res2, err := stmt2.Exec(3)
	if err != nil {
		log.Fatal(err)
	}

	rows2, _ := res2.RowsAffected()
	fmt.Println(rows2)
}
