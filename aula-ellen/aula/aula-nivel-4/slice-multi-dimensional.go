package main

import "fmt"

func main() {

	ss := [][]int{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{7, 8, 9, 10, 11, 12},
		[]int{13, 14, 15, 16, 17, 18},
	}
	fmt.Println(ss[2][4])
}
