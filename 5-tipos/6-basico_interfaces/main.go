package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - $%.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var qualquerCoisa imprimivel = pessoa{"Rafael", "Dantas"}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	qualquerCoisa = produto{"Lapis", 1.99}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	imprimir(produto{"Caneta", 5.99})
}
