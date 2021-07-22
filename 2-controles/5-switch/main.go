package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func notaParaConceitoSimplificada(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "texto"
	case func():
		return "função"
	default:
		return "desconhecido"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11))
	fmt.Println(notaParaConceito(-1))
	fmt.Println(notaParaConceito(0))
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(10.1))

	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	fmt.Println(notaParaConceitoSimplificada(9.8))
	fmt.Println(notaParaConceitoSimplificada(6.9))
	fmt.Println(notaParaConceitoSimplificada(2.1))
	fmt.Println(notaParaConceitoSimplificada(11))
	fmt.Println(notaParaConceitoSimplificada(-1))
	fmt.Println(notaParaConceitoSimplificada(0))
	fmt.Println(notaParaConceitoSimplificada(10))
	fmt.Println(notaParaConceitoSimplificada(10.1))

	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
