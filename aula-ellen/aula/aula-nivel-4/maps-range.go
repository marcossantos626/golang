package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}
	fmt.Println(qualquercoisa)
	for i, v := range qualquercoisa {
		fmt.Println(i, v)

	}
}
