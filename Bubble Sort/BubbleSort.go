package main

import (
	"fmt"
)

// Swap intercambia el elemento en el índice i con el elemento en i+1
func Swap(numbers []int, i int) {
	// En Go podemos intercambiar valores en una sola línea
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

// BubbleSort ordena la rebanada (slice) in-place de menor a mayor
func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		// El último i elementos ya están en su lugar
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	// Definimos un slice para guardar hasta 10 números
	var input int
	numbers := make([]int, 0, 10)

	fmt.Println("Ingrese hasta 10 números enteros (escriba un número no válido o alcance 10 para terminar):")

	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		_, err := fmt.Scan(&input)
		if err != nil {
			// Si el usuario ingresa algo que no es un número, dejamos de pedir
			break
		}
		numbers = append(numbers, input)
	}

	fmt.Println("\nLista original:", numbers)

	// Llamamos a la función de ordenamiento
	BubbleSort(numbers)

	fmt.Println("Lista ordenada:", numbers)
}