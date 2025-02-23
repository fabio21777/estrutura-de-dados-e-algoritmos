package main

import (
	"fmt"
)

func pessoa_e_vendedor(nome string) bool {
	return nome[len(nome)-1] == 'm'
}

var grafo = map[string][]string{
	"voce":   {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"peggy"},
	"claire": {"thom", "jonny"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func pesquisa(nome string) bool {
	fila := []string{nome}
	verificados := make(map[string]bool)

	for len(fila) > 0 {
		pessoa := fila[0]
		fila = fila[1:]

		if !verificados[pessoa] {
			fmt.Println("Verificando " + pessoa)
			if pessoa_e_vendedor(pessoa) {
				fmt.Println(pessoa + " é um vendedor de manga!")
				return true
			}
			verificados[pessoa] = true
			fila = append(fila, grafo[pessoa]...)
		}
	}

	return false
}

func main() {
	if !pesquisa("voce") {
		fmt.Println("Nenhum vendedor de manga encontrado.")
	}
}
