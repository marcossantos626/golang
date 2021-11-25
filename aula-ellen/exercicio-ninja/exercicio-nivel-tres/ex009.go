//Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	esporteFavorito := "vôlei"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Neymar é um jogador de Futebol)")
	case "vôlei":
		fmt.Println("Bernardo Rocha de Rezende, conhecido como Bernardinho, é um ex-jogador, treinador de voleibol, economista, e empresário brasileiro.")
	case "ping-pong":
		fmt.Println("No ping-pong, você compra a raquete e ela já vem pronta para o jogo.\nAlgumas possuem até EVA em um dos lados.\nPorém, no tênis de mesa você compra a raquete, a borracha e a cola para, então, montar sua raquete.\nAlém do mais, a borracha precisa ser certificada pela Federação Internacional de Tênis de Mesa (ITTF).")
	}
}
