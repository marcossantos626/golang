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
