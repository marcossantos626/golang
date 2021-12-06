package main

import "fmt"

func main() {
	slice := []string{"Banana", "Maça", "Jaca", "Pêssego"}

	slice = append(slice, "Melancia")

	for indice, valor := range slice {
		fmt.Println("No Índice", indice, "temos o valor:", valor)
	}

}

func exemplo() {
	slice := []int{20, 21, 22, 23}
	total := 0
	for _, valor := range slice {
		total += valor
	}
	fmt.Println(" O valor total é:", total)
}
