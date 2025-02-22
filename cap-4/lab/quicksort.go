package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quicksort(lista []int) []int {
	// condição de parada
	if len(lista) < 2 { // condição base quando o array é igual a 1 ou zero
		return lista
	} else { // condição recursiva
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		pivoIndex := r.Intn(len(lista))
		pivo := lista[pivoIndex]
		menores := []int{}
		maiores := []int{}

		for i, v := range lista {
			if i == pivoIndex {
				continue
			}
			if v <= pivo {
				menores = append(menores, v)
			} else {
				maiores = append(maiores, v)
			}
		}

		return append(append(quicksort(menores), pivo), quicksort(maiores)...)
	}
}

func main() {
	lista := []int{10, 5, 2, 3}
	fmt.Println(quicksort(lista))
}
