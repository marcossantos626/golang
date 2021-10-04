package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.78,
			"Guga Pereira":   54321.87,
		},
		"J": {
			"José João": 98765.09,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
