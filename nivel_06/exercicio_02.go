package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(slice...))
	fmt.Println(sum2(slice))
}

func sum(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func sum2(numbers []int) int {
	return sum(numbers...)
}