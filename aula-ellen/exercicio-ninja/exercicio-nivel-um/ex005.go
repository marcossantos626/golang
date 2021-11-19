package main

import "fmt"

type deuses int

var x deuses
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println("<----- Espaço ----->")

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
