package main

import "fmt"

func main() {
	soma := retornaFunc()

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(soma(numeros...))
}

func retornaFunc() func(numeros ...int) int {
	return func(numeros ...int) int {
		soma := 0
		for _, valor := range numeros {
			soma += valor
		}
		return soma
	}
}