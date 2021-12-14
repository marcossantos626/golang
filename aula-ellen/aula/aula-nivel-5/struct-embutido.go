package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {

	pessoa1 := pessoa{
		nome:  "Zeus",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Netuno",
			idade: 37,
		},
		titulo:  "QA",
		salario: 7985.98,
	}
	pessoa3 := profissional{
		pessoa: pessoa{
			nome:  "Minerva",
			idade: 39,
		},
		titulo:  "RH",
		salario: 3885.98,
	}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
}
