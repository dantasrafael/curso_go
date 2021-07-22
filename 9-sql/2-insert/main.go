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

	stmt, err := db.Prepare(`INSERT INTO usuarios(nome) VALUES($1) RETURNING id`)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec("Pedro")
	stmt.Exec("João")

	res, _ := stmt.Exec("Rafael")
	fmt.Println("Res:", res)
	fmt.Println(res.LastInsertId()) // Não suportado por este driver
	fmt.Println(res.RowsAffected())

	var userId int
	db.QueryRow(`INSERT INTO usuarios(nome) VALUES($1) RETURNING id`, "Maria").Scan(&userId)
	fmt.Println("Novo id:", userId)

	err = db.QueryRow(`INSERT INTO usuarios(id, nome) VALUES(1, 'Tiago')`).Scan()
	if err != nil {
		fmt.Println("Ocorreu um erro!!!")
		log.Fatal(err)
	}
}
