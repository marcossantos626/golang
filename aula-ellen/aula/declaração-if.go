package main

import "fmt"

var x = 90

func main() {
	if x > 100 {
		fmt.Println("Esse valor", x, "não é menor que 100")
	}
	if x < 100 {
		fmt.Println("Esse valor", x, "é menor que 100")
	}
}
