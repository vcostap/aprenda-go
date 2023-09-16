package main

import "fmt"

func main() {
	a := 20
	b := a << 1

	fmt.Printf("decimal(a) = %d / binario(a) = %b / hexadecial(a) %#x\n", a, a, a)
	fmt.Printf("decimal(b) = %d / binario(b) = %b / hexadecial(b) %#x\n", b, b, b)
}