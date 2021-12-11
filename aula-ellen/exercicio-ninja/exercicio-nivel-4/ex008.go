//Crie um map com key tipo string e value tipo []string.
//Key deve conter nomes no formato sobrenome_nome
//Value deve conter os hobbies favoritos da pessoa
//Demonstre todos esses valores e seus índices.

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
	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, h)
		}
	}
	fmt.Println(mepe)

}
