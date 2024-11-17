package main

import "fmt"

func buscaBinaria(lista []int, item int) int {
	limiteInferior := 0
	limiteSuperior := len(lista) - 1
	index := 1
	for limiteInferior <= limiteSuperior {
		meio := (limiteInferior + limiteSuperior) / 2
		chute := lista[meio]
		fmt.Println(fmt.Sprintf("meio %d", meio))
		fmt.Println(fmt.Sprintf("chute %d", chute))
		if item == chute {
			fmt.Println(fmt.Sprintf("Iterações %d", index))
			return meio
		}
		if chute > item {
			limiteSuperior = meio - 1
		}

		if chute < item {
			limiteInferior = meio + 1
		}
		index++
	}
	return -1

}

func main() {
	lista := []int{1, 3, 5, 7, 9}
	fmt.Println("index que contem o valor:", buscaBinaria(lista, 9))
	fmt.Println(buscaBinaria(lista, -1))
}
