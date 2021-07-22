package main

import "fmt"

func main() {
	s := make([]int, 10) // slice de inteiros com 10 elementos
	s[9] = 1
	fmt.Println(s)

	s = make([]int, 10, 20) // slice de inteiros com 10 elementos e array interno com 20 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
