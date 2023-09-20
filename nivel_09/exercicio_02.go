package main

import "fmt"

type pessoa struct {
	fala string
}

type humanos interface {
	falar()
}

func main() {
	pessoa := pessoa{"Eu sou uma pessoa"}
	fmt.Println("Usando o método falar(), pois o tipo pessoa possui esse método.")
	pessoa.falar() // atalho para a linha debaixo.
	(&pessoa).falar()
	fmt.Println("----------------")
	fmt.Println("Usando a função dizerAlgumaCoisa(), pois ela recebe uma interface humano como parâmetro.")
	dizerAlgumaCoisa(&pessoa)
	//dizerAlgumaCoisa(pessoa) <- não funciona, pois a função falar() é implementada para um ponteiro para pessoa e não uma pessoa.
}

func (p *pessoa) falar() {
	fmt.Println(p.fala)
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}
