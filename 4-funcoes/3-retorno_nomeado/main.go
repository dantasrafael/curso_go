package main

import "fmt"

// func troca(p1, p2 int) (segundo int, primeiro int) {
// 	return p2, p1
// }

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	x, y := troca(1, 2)
	fmt.Println(x, y)

	segundo, primeiro := troca(3, 4)
	fmt.Println(segundo, primeiro)
}
