package main

import (
	"fmt"
	"time"
)

func main() {
	/* Una condición de carrera ocurre cuando dos o más goroutines 
	   intentan acceder y modificar la misma variable compartida 
	   al mismo tiempo.
	*/
	variableCompartida := 0

	// Lanzamos la primera goroutine para incrementar el valor
	go func() {
		for i := 0; i < 100; i++ {
			valor := variableCompartida
			// Simulamos un pequeño retraso para forzar la interrupción
			time.Sleep(time.Millisecond) 
			variableCompartida = valor + 1
		}
	}()

	// Lanzamos la segunda goroutine con el mismo propósito
	go func() {
		for i := 0; i < 100; i++ {
			valor := variableCompartida
			time.Sleep(time.Millisecond)
			variableCompartida = valor + 1
		}
	}()

	// Esperamos a que terminen (tiempo suficiente para el ejemplo)
	time.Sleep(2 * time.Second)

	/* Si no hubiera condición de carrera, el resultado debería ser 200.
	   Debido a la condición de carrera, el resultado será menor e impredecible.
	*/
	fmt.Printf("Valor final de la variable: %d\n", variableCompartida)
}