package main

import "fmt"

func main() {
	sabores := []string{"Pepperoni", "Mozzarela", "Abacaxi", "Quatro-queijo", "Marguerita"}
	fatia := sabores[:]
	fmt.Println(fatia)
	//suando for

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])

	}
}
