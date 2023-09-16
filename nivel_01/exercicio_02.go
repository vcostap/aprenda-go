package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Printf("%v -> %T, %v -> %T, %v -> %T\n", x, x, y, y, z, z)
	// o valor atribuido às variáveis pelo compilador se chama "valor zero"

}