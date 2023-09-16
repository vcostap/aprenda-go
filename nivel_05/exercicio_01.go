package main

import "fmt"

type pessoas struct {
	nome          string
	sobrenome     string
	sabores_pizza []string
}

func main() {
	pessoas := []pessoas{
		{
			nome:          "Vinícius",
			sobrenome:     "Costa",
			sabores_pizza: []string{"Pepperoni", "Lombinho com catupiry"},
		},
		{"Paula", "Porto", []string{"Calabresa"}},
	}

	for _, pessoa := range pessoas {
		fmt.Println("Nome:", pessoa.nome)
		fmt.Println("Sobrenome:", pessoa.sobrenome)
		fmt.Println("Sabores de pizza:")
		for _, sabor := range pessoa.sabores_pizza {
			fmt.Println("\t", sabor)
		}
		fmt.Println("-----------------")
	}
}