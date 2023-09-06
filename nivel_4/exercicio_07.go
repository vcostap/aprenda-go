package main

import "fmt"

func main() {
	pessoas := [][]string{
		{
			"Vinícius",
			"Costa",
			"Futebol",
		},
		{
			"Paula",
			"Porto",
			"Trabalho",
		},
		{
			"Gabriel",
			"Costa",
			"Crossfit",
		},
	}

	for pessoa := range pessoas {
		for _, value := range pessoas[pessoa] {
			fmt.Println(value)
		}
		fmt.Println("--------")
	}

}