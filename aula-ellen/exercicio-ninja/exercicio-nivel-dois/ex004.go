//Crie um programa que:
//Atribua um valor int a uma variável
//Demonstre este valor em decimal, binário e hexadecimal
//Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import "fmt"

func main() {
	a := 200
	fmt.Println("A = 200")
	fmt.Printf("binário = %b\nDecimal = %d\nhexadecimal = %#x\n", a, a, a)

	fmt.Println(" ")
	fmt.Println("------Espaço------")
	fmt.Println(" ")

	fmt.Println("B = A << 1")
	b := a << 1
	fmt.Printf("binário = %b\nDecimal = %d\nhexadecimal = %#x\n", b, b, b)

	fmt.Println(" ")
	fmt.Println("------Espaço------")
	fmt.Println(" ")

	fmt.Println("C = A >> 1")
	c := a >> 1
	fmt.Printf("binário = %b\nDecimal = %d\nhexadecimal = %#x", c, c, c)
}
