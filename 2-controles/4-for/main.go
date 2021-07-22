package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 20; j += 2 {
		fmt.Printf("%d ", j)
	}

	fmt.Println("\nMisturando...")
	for k := 1; k <= 10; k++ {
		var tipo = "impar"
		if k%2 == 0 {
			tipo = "par"
		}

		fmt.Printf("| %d é %s ", k, tipo)
	}

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
