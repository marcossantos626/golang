//Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//"Nome", "Sobrenome", "Hobby favorito"
//Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {
	ss := [][]string{
		[]string{
			"Jupiter",
			"Maximus",
			"Estudar",
		},
		[]string{
			"Aglaia",
			"Maximus",
			"Pintar",
		},
		[]string{
			"Minerva",
			"Atena",
			"Ler",
		},
	}
	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}
}
