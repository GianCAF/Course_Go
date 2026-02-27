package main
import (
	"fmt"
	"sort"
	"strconv"
)
func main() {
	// Creamos una rebanada vacía de enteros
	rebanada := make([]int, 0)
	for {
		var entrada string
		fmt.Scan(&entrada)
		// Si el usuario escribe 'X' o algo que no es número, el curso suele terminar
		// Pero aquí simplemente convertimos la entrada a entero
		numero, err := strconv.Atoi(entrada)
		if err != nil {
			break
		}
		// Añadimos el número a la rebanada
		rebanada = append(rebanada, numero)
		// Ordenamos la rebanada de menor a mayor
		sort.Ints(rebanada)
		// Imprimimos la rebanada ordenada
		fmt.Println(rebanada)
	}
}