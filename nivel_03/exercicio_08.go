package main

import (
	"fmt"
)

func main() {
	a := 13

	switch {
	case a < 5:
		fmt.Println("a é menor que 5")
	case a < 10 && a >= 5:
		fmt.Println("a é maior ou igual a cinco e menor que 10")
	default:
		fmt.Println("a é maior ou igual a 10")
	}
}