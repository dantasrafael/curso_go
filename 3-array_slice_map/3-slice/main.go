package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Slice não é um array. Slide define um pedaço de um array com o uso de ponteiro
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // a partir do índice 1 até o 3, não incluíndo o 3 elemento
	fmt.Println(a2, s2)

	a2[1] = 10
	fmt.Println(a2, s2)
	fmt.Println(&a2[1])
	fmt.Println(&s2[1])

	s3 := a2[:2] // a partir do ínicio do array até a posição 2, não incluíndo o 2 elemento
	fmt.Println(a2, s3)

	a2[2] = 20
	fmt.Println(a2, s3)

	s4 := s2[:1] // slice de slice
	fmt.Println(a2, s2, s4)
}
