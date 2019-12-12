package main

import "fmt"

func notaParaConceito(n float64) {
	switch {
	case n >= 9 && n <= 10:
		fmt.Println("A")
	case n >= 8 && n < 9:
		fmt.Println("B")
	case n >= 5 && n < 8:
		fmt.Println("C")
	case n >= 3 && n < 5:
		fmt.Println("D")
	case n >= 0 && n < 3:
		fmt.Println("E")
	default:
		fmt.Println("Nota Inválida!")
	}
}

func main() {
	notaParaConceito(9.5)
	notaParaConceito(8.5)
	notaParaConceito(6)
	notaParaConceito(4.5)
	notaParaConceito(0.5)
	notaParaConceito(-10)
}
