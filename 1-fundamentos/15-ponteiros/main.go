package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço de i
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiro
	// p++

	fmt.Println(p, *p, i, &i)

}
