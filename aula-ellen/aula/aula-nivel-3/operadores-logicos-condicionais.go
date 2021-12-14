package main

import "fmt"

func main() {
	x := 12
	if x == 2 || x == 3 {
		fmt.Println("X é igual a DOIS ou TRÊS")
	} else {
		fmt.Println("X NÃO é igual a DOIS ou TRÊS")
	}
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é multiplo de DOIS e TRÊS")
	} else {
		fmt.Println("X NÃO é multiplo de DOIS e TRÊS")
	}
}
