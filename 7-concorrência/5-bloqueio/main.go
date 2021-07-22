package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	fmt.Println("Após sleep de 1s")
	c <- 1 // operação bloqueante
	fmt.Println("Depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)
	fmt.Println("Antes de ler")
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
}
