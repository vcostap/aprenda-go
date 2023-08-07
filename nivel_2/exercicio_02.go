package main

import "fmt"

func main() {
	a, b, c := 1, 10, 10
	igual := a == b
	diferente := a != b
	menor_igual := a <= b
	maior_igual := a >= b
	menor := b < c
	maior := b > a

	fmt.Printf("a == b? %v\n", igual)
	fmt.Printf("a != b? %v\n", diferente)
	fmt.Printf("a <= b? %v\n", menor_igual)
	fmt.Printf("a >= b? %v\n", maior_igual)
	fmt.Printf("b < c? %v\n", menor)
	fmt.Printf("b > a? %v\n", maior)
}