//Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("binário = %b\nDecimal = %d\nhexadecimal = %#x", x, x, x)
}
