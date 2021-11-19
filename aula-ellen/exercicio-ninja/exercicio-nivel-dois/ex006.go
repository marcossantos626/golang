package main

import "fmt"

const (
	_ = 2022 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
