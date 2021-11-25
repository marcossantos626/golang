package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 3 {
			//faz isso quando número não é par
			continue
		}
		fmt.Println(i)
	}
}
