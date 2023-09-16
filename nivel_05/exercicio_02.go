package main

import "fmt"

type pessoas struct {
	nome          string
	sobrenome     string
	sabores_pizza []string
}

func main() {
	pessoas := map[string]pessoas{
		"Costa": {
			nome:          "Vinícius",
			sobrenome:     "Costa",
			sabores_pizza: []string{"Pepperoni", "Lombinho com catupiry"},
		},
		"Porto": {"Paula", "Porto", []string{"Calabresa"}},
	}

	for key, pessoa := range pessoas {
		fmt.Println("===", key, "===")
		fmt.Println("Nome:", pessoa.nome)
		fmt.Println("Sobrenome:", pessoa.sobrenome)
		fmt.Println("Sabores de pizza:")
		for _, sabor := range pessoa.sabores_pizza {
			fmt.Println("\t", sabor)
		}
		fmt.Println("-----------------")
	}
}