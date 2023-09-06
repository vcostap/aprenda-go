package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"Costa_Vinícius": {"Futebol", "Valorant"},
		"Porto_Paula":    {"Trabalho", "Viajar"},
		"Costa_Gabriel":  {"Crossfit", "Pedalar"},
	}

	pessoas["Costa_Regina"] = []string{"Rezar", "Ir na feira hippie"}

	for key, hobbies := range pessoas {
		fmt.Printf("name -> %v\n\thobbies:", key)
		for i, hobby := range hobbies {
			fmt.Printf("\n\t\t%v - %v", i, hobby)
		}
		fmt.Printf("\n\n")
	}

	fmt.Printf("%T\n", pessoas)

}