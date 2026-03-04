package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn genera una función que calcula el desplazamiento
// Recibe: aceleración (a), velocidad inicial (vo) y desplazamiento inicial (so)
// Devuelve: una función que acepta el tiempo (t) y devuelve el desplazamiento (s)
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	// Retornamos una función anónima (clausura) que "recuerda" los valores de a, vo y so
	return func(t float64) float64 {
		// Fórmula: s = 0.5 * a * t^2 + vo * t + so
		return (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Println("--- Calculador de Desplazamiento ---")
	
	// 1. Pedir valores iniciales al usuario
	fmt.Print("Ingrese la aceleración (a): ")
	fmt.Scan(&a)
	fmt.Print("Ingrese la velocidad inicial (vo): ")
	fmt.Scan(&vo)
	fmt.Print("Ingrese el desplazamiento inicial (so): ")
	fmt.Scan(&so)

	// 2. Generar la función de desplazamiento usando los valores ingresados
	fn := GenDisplaceFn(a, vo, so)

	// 3. Pedir el tiempo (t)
	fmt.Print("Ingrese el tiempo (t) para calcular el desplazamiento: ")
	fmt.Scan(&t)

	// 4. Calcular y mostrar el resultado usando la función generada
	fmt.Printf("\nEl desplazamiento final tras %.2f segundos es: %.2f\n", t, fn(t))
}