package main

import "fmt"

func main() {
	estados := make([]string, 5, 27)
	resposta := []string{"Acre", "Alagoas", "Amapá", "Amazonas"}

	_ = copy(estados, resposta)

	fmt.Printf("len -> %v\ncap -> %v\n", len(estados), cap(estados))

	for i := range estados {
		fmt.Println(estados[i], "\n", "--------")
	}
}