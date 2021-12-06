package main

import "fmt"

func main() {
	slice := []int{20, 21, 22, 23, 1, 13}
	total := 0
	for _, valor := range slice {
		total += valor
	}
	fmt.Println(" O valor total é:", total)
}
