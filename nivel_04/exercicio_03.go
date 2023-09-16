package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(x[:3])
	fmt.Println(x[4:])
	fmt.Println(x[1:7])
	fmt.Println(x[2:9])
	fmt.Println(x[2 : len(x)-1])
}