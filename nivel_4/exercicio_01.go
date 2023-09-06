package main

import "fmt"

func main() {
	x := [5]int{}

	x[0], x[1], x[2], x[3], x[4] = 10, 20, 30, 40, 50

	for i := range x {
		fmt.Println(x[i])
	}

	fmt.Printf("x é do tipo: %T\n", x)
}