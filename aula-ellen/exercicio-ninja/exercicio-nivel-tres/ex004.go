//Crie um loop utilizando a sintaxe: for {}
//Utilize-o para demonstrar os anos desde que você nasceu.
package main

import "fmt"

func main() {
	ap := 1996
	af := 2060
	for {
		if ap > af {
			break
		}
		fmt.Println(ap)
		ap++

	}
}
