package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // [...] compilador que vai contar!

	// indice e valores
	fmt.Println("índice e valores")
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// somente os valores
	fmt.Println("\nsomente valores")
	for _, num := range numeros {
		fmt.Println(num)
	}

	// somente o índice
	fmt.Println("\nsomente índice")
	for index := range numeros {
		fmt.Println(index)
	}
}
