package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	eu := pessoa{
		nome:      "Vinícius",
		sobrenome: "Costa",
		idade:     30,
	}

	eu.mostreDados()
}

func (p pessoa) mostreDados() {
	fmt.Printf("O nome completo da pessoa é: %s %s.\nEla possui %d anos.\n", p.nome, p.sobrenome, p.idade)
}