package main

import "fmt"

func main() {
	x, y, z := 42, "James Bond", true
	fmt.Printf("%v -> %T, %v -> %T, %v -> %T\n", x, x, y, y, z, z)
	fmt.Println("=================")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}