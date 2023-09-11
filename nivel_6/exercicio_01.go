package main

import "fmt"

func main() {
	x := print5()
	y, z := hellow()

	fmt.Printf("%s! Tenho %d ovos. Logo preciso de %d pães no café da manhã.\n", z, x, y)
}

func print5() int {
	return 5
}

func hellow() (int, string) {
	return 10, "Bom dia"
}