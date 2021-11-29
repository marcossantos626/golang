package main

import "fmt"

func main() {
	if x := 99; x > 100 {
		fmt.Println("O valor inserido é MAIOR que CEM")
	} else if x < 100 {
		fmt.Println("O valor inserido é MENOR que CEM")
	} else {
		fmt.Println("O valor inserido NÃO é MAIOR nem MENOR que CEM")
	}
}
