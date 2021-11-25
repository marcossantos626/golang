package main

import (
	"fmt"
)

var x interface{}

func main() {
	x = 666
	switch x.(type) {
	case int:
		fmt.Printf("É um INT")
	case bool:
		fmt.Println("É um BOOL")
	case string:
		fmt.Println("É um STRING")
	case float64:
		fmt.Println("É um FLOAT")
	}
}
