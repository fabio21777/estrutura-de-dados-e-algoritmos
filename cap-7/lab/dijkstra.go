package main

import (
	"fmt"
	"math"
)

// Definição do grafo
var grafo = map[string]map[string]int{
	"A": {"B": 2, "C": 4},
	"B": {"C": 1, "D": 7},
	"C": {"D": 3, "E": 5},
	"D": {"E": 1, "F": 4},
	"E": {"F": 2},
	"F": {},
}

// Inicialização das distâncias
var distancia = map[string]int{
	"A": 0,
	"B": math.MaxInt64,
	"C": math.MaxInt64,
	"D": math.MaxInt64,
	"E": math.MaxInt64,
	"F": math.MaxInt64,
}

var visitados = make(map[string]bool)

func menor_distancia(nome string) {
	visitados[nome] = true
	for vizinho, peso := range grafo[nome] {
		if distancia[vizinho] > distancia[nome]+peso {
			distancia[vizinho] = distancia[nome] + peso
		}
	}
}

func main() {
	for len(visitados) < len(grafo) {
		// Encontrar o vértice com a menor distância que ainda não foi visitado
		menor := math.MaxInt64
		vertice := ""

		for k, v := range distancia {
			if !visitados[k] && v < menor {
				menor = v
				vertice = k
			}
		}

		// Se nenhum vértice foi encontrado, significa que o restante é inacessível
		if vertice == "" {
			break
		}

		menor_distancia(vertice)
	}

	// Exibir distâncias mínimas
	for k, v := range distancia {
		fmt.Printf("Distância de A até %s: %d\n", k, v)
	}
}
