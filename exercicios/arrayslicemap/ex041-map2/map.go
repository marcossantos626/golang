package main

import "fmt"

func main() {
	funcsEsalarios := map[string]float64{
		"Zeus":    1234.56,
		"Júpiter": 7890.12,
		"Marcos":  3212.45,
	}

	funcsEsalarios["Hades"] = 4321.87
	delete(funcsEsalarios, "Inexistente")

	for nome, salario := range funcsEsalarios {
		fmt.Println(nome, salario)
	}
}
