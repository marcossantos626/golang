package main

import "fmt"

func main() {
	quemvaisextar := "dinheiro"
	switch quemvaisextar {
	case "Jupiter":
		fmt.Println("Quem vai sextar é o Jupiter")
	case "Plutão":
		fmt.Println("Quem vai sextar é o Plutão")
	case "Netuno":
		fmt.Println("Quem vai sextar é o Netuno")
	case "Aglaia":
		fmt.Println("Quem vai sextar é o Aglaia")
	case "Minerva":
		fmt.Println("Quem vai sextar é a Minerva")
	case "Juno":
		fmt.Println("Quem vai sextar é a Juno")
	case "Diana":
		fmt.Println("Quem vai sextar é a Diana")
	case "Vesta":
		fmt.Println("Quem vai sextar é a Vesta")
	default:
		fmt.Println("Ninguem tem dinheiro")

	}
}
