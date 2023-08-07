package main

import "fmt"

func main() {
	a, b, c := 1, 10, 20
	fmt.Printf("decimal(a) -> %v\nbinário(a) -> %b\nhexadecimal(a) -> %#x\n", a, a, a)
	fmt.Printf("decimal(b) -> %v\nbinário(b) -> %b\nhexadecimal(b) -> %#x\n", b, b, b)
	fmt.Printf("decimal(c) -> %v\nbinário(c) -> %b\nhexadecimal(c) -> %#x\n", c, c, c)
}