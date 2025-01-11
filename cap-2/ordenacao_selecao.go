package main

func ordenacaoSelecao(lista []int) []int {
	novoArray := []int{}
	tamanholista := len(lista)

	// enquanto o tamanho da lista for maior que 0 devo continuar

	for i := 0; i < tamanholista; i++ {
		menor, index := encontraMenor(lista)
		novoArray = append(novoArray, menor)
		lista = remove(index, lista)
	}

	return novoArray
}

func encontraMenor(lista []int) (int, int) {
	menor := lista[0]
	index := 0
	for i := 1; i < len(lista); i++ {
		if lista[i] < menor {
			menor = lista[i]
			index = i
		}
	}

	return menor, index
}

func remove(index int, lista []int) []int {
	return append(lista[:index], lista[index+1:]...)
}

func main() {
	lista := []int{5, 3, 6, 2, 10}
	novoArray := ordenacaoSelecao(lista)

	for _, v := range novoArray {
		println(v)
	}
}
