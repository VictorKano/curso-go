package main

import "fmt"

func main() {
	//var aprovados map[int]string  //Mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95134657399] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95134657399])
	delete(aprovados, 95134657399)

	fmt.Println(aprovados)
}
