package main

import (
	"fmt"
)

func main() {
	if a := 13; a < 5 {
		fmt.Println("a é menor que 5")
	} else if a < 10 && a >= 5 {
		fmt.Println("a é maior ou igual a cinco e menor que 10")
	} else {
		fmt.Println("a é maior ou igual a 10")
	}
}