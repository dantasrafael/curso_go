package main

import "fmt"

func inc1(n int) {
	n++
}

// Cuidado: essa função se tornou "não pura",
// pois altera um valor que não faz parte de seu contexto
func inc2(n *int) { // revisão: um ponteiro é representado por um *
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor
	fmt.Println(n) // não altera o valor de N

	// revisão: & é usado para obter o endereço da variável
	inc2(&n)       // por referência
	fmt.Println(n) // altera o valor de N
}
