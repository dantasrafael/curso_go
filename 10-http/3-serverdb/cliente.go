package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Usuario :)
type Usuario struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type UsuarioSave struct {
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega a função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	case r.Method == "DELETE" && id > 0:
		excluir(w, r, id)
	case r.Method == "POST":
		inserir(w, r)
	case r.Method == "PUT" && id > 0:
		atualizar(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func getConexaoDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=curso_golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db := getConexaoDB()
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = $1", id).Scan(&u.Id, &u.Nome)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := getConexaoDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, nome FROM usuarios ORDER BY id")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ocorreu um erro no servidor, desculpa... :(")
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.Id, &u.Nome)
		usuarios = append(usuarios, u)
	}
	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func excluir(w http.ResponseWriter, r *http.Request, id int) {
	db := getConexaoDB()
	defer db.Close()

	stmt, _ := db.Prepare(`DELETE FROM usuarios WHERE id = $1`)

	_, err := stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ocorreu um erro ao tentar excluir, desculpa... :(")
	}
	w.WriteHeader(http.StatusNoContent)
}

func inserir(w http.ResponseWriter, r *http.Request) {
	db := getConexaoDB()
	defer db.Close()

	var dto UsuarioSave
	json.NewDecoder(r.Body).Decode(&dto)

	u := Usuario{Nome: dto.Nome}
	db.QueryRow(`INSERT INTO usuarios(nome) VALUES($1) RETURNING ID`, dto.Nome).Scan(&u.Id)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(json))
}

func atualizar(w http.ResponseWriter, r *http.Request, id int) {
	db := getConexaoDB()
	defer db.Close()

	var dto UsuarioSave
	json.NewDecoder(r.Body).Decode(&dto)

	u := Usuario{Id: id, Nome: dto.Nome}
	db.QueryRow(`UPDATE usuarios SET nome = $1 WHERE id = $2`, dto.Nome, id)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, string(json))
}
