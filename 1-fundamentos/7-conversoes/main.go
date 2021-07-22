package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(rune(123)))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para bool
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseBool("false")
	c, _ := strconv.ParseBool("0")
	d, _ := strconv.ParseBool("1")
	fmt.Println(a, b, c, d)
}
