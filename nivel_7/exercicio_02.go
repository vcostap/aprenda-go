package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	pessoa := pessoa{"Vinícius", "Costa", 30}
	fmt.Println(pessoa)

	mudeMe(&pessoa)
	fmt.Println(pessoa)
}

func mudeMe(pessoa *pessoa) {
	// A maneira correta de fazer dereference de um valor em uma struct é: (*struct).field
	// Existe uma forma abreviada, que é struct.field

	// Jeito abreviado
	pessoa.sobrenome = "Sobrenome"
	// Jeito correto
	(*pessoa).nome = "Nome"
	// Jeito abreviado
	pessoa.idade = 99
}