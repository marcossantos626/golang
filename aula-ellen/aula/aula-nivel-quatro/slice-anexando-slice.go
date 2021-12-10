package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	fmt.Println("Antiga   ", umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println("Atual    ", umaslice)

	doisslice := []int{9, 10, 11, 12}
	umaslice = append(umaslice, doisslice...)
	fmt.Println("Juntando ", umaslice)

}
