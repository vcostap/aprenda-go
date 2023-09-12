package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	soma := pegaNumerosEFazAlgo(soma, numeros...)

	fmt.Println(soma)
}

func pegaNumerosEFazAlgo(f func(numeros ...int) int, numeros ...int) int {
	fmt.Println(numeros)

	return f(numeros...)
}

func soma(numeros ...int) int {
	soma := 0
	for _, valor := range numeros {
		soma += valor
	}
	return soma
}