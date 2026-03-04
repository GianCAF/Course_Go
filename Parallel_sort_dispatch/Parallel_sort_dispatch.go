package main

import (
	"fmt"
	"sort"
	"sync"
)

// Función para ordenar una submatriz e imprimirla
func sortSubarray(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Cada goroutine imprime el conjunto que le tocó
	fmt.Println("Goroutine ordenando submatriz:", arr)
	sort.Ints(arr)
}

func main() {
	var n int
	fmt.Print("¿Cuántos números desea ingresar? ")
	fmt.Scan(&n)

	slice := make([]int, n)
	fmt.Println("Ingrese los números uno por uno:")
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	// Calculamos el tamaño de cada una de las 4 partes
	size := (len(slice) + 3) / 4
	var wg sync.WaitGroup

	// Particionamos el array en 4 partes aproximadamente iguales
	for i := 0; i < 4; i++ {
		start := i * size
		end := (i + 1) * size
		if start > len(slice) {
			start = len(slice)
		}
		if end > len(slice) {
			end = len(slice)
		}

		wg.Add(1)
		// Cada partición es ordenada por una goroutine diferente
		go sortSubarray(slice[start:end], &wg)
	}

	wg.Wait()

	// La goroutina principal fusiona las 4 submatrices
	sort.Ints(slice)

	// Imprimir la lista final ordenada
	fmt.Println("\nLista final ordenada por la goroutine principal:")
	fmt.Println(slice)
}