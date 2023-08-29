package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("futebol")
	case "natação":
		fmt.Println("natação")
	case "vôlei":
		fmt.Println("vôlei")
	case "crossfit":
		fmt.Println("crossfit")
	default:
		fmt.Println("que esporte o que, mané?!")
	}
}