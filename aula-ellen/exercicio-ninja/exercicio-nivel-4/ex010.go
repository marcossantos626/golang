//Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {
	mepe := map[string][]string{
		"Grego": []string{
			"Origem dos Deuses", "Zeus", "Autoritário",
		},
		"Roma": []string{
			"Novos Deuses", "Jupiter", "O Melhor e o Maior",
		},
	}
	mepe["Nova Ordem"] = []string{"Nova Rainha", "Aglaia", "Esposa de Zeus/Jupiter"}
	delete(mepe, "Grego")
	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, h)
		}
	}

}
