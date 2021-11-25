package main

import "fmt"

func main() {
	for h := 0; h <= 12; h++ {
		fmt.Println("Hora:", h)
		for m := 0; m <= 60; m++ {
			fmt.Print(" ", m)
		}
		fmt.Print("\n")
	}
}
