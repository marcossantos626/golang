package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := cliente{
		nome:      "Jupiter ",
		sobrenome: "Optimus Maximus",
		fumante:   false,
	}
	cliente2 := cliente{"Aglaia", "Optimus Maximus", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}
