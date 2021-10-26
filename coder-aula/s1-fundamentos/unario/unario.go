package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // X += 1 ou x = x + 1
	fmt.Println(x)

	y-- // X -= 1 ou x = y - 1
	fmt.Println(y)

	//fmt.Println(++x == y--) não existe
}
