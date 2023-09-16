package main

import "fmt"

func main() {
	anonimo := struct {
		pessoa           map[string]string
		cidadesVisitadas []string
	}{
		pessoa: map[string]string{
			"nome":      "Vinícius",
			"sobrenome": "Costa",
		},
		cidadesVisitadas: []string{"João Pessoa", "Sydney", "Melbourne", "São Miguel do Gostoso"},
	}

	fmt.Println(anonimo)
}