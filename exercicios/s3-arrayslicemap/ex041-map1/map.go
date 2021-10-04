package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//Mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123455678] = "Zeus"
	aprovados[765432234] = "Júpiter"
	aprovados[328773737] = "Marcos"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[328773737])
	delete(aprovados, 328773737)
	fmt.Println(aprovados[328773737])
}
