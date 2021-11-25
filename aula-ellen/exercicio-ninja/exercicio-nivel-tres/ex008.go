//Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {
	jupiter := 3

	switch {
	case jupiter == 0:
		fmt.Println("Você é lindo(a)")
	case jupiter == 1:
		fmt.Println("Jogue na mega")
	case jupiter == 3:
		fmt.Println("Hora do Lazer")
	}
}
