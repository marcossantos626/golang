//Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//Do quinto ao último item do slice (incluindo o último item!)
//Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {
	//               0   1   2   3   4   5   6   7   8   9
	slice := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println("Slice principal:                            ", slice)
	fmt.Println("Do primeiro ao terceiro item:               ", slice[:3])
	fmt.Println("Do quinto ao último item:                   ", slice[4:])
	fmt.Println("Do segundo ao sétimo item:                  ", slice[1:7])
	fmt.Println("Do terceiro ao penúltimo item:              ", slice[2:9])
	fmt.Println("Do terceiro ao penúltimo item (usando len): ", slice[2:len(slice)-1])
}
