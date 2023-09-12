package main

import "fmt"

func main() {
	magic := func(num int) {
		fmt.Println("O número é:", num)
	}

	magic(7)
}