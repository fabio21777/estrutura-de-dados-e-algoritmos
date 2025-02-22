package main

import (
	"fmt"
)

func fat(n int) int {
	if n == 0 { // caso base
		fmt.Println("Cheguei no caso base")
		return 1
	}
	fmt.Println("Calculando fat(", n, ") - caso recursivo")
	return n * fat(n-1)
}

func main() {
	fmt.Println(fat(5))
}
