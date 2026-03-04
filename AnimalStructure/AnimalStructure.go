package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal representa la estructura base con tres campos de tipo cadena
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat imprime la comida del animal
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move imprime el método de locomoción
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak imprime el sonido del animal
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Inicializamos el mapa con los datos de la tabla proporcionada
	animals := map[string]Animal{
		"vaca":     {food: "hierba", locomotion: "caminar", noise: "mugido"},
		"pájaro":   {food: "gusanos", locomotion: "volar", noise: "pío"},
		"serpiente": {food: "ratones", locomotion: "deslizarse", noise: "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Interfaz de información de animales")
	fmt.Println("Instrucciones: Escribe el animal y la propiedad (ejemplo: vaca comer)")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		// Separamos la línea en palabras
		input := strings.Fields(strings.ToLower(scanner.Text()))

		// Validamos que se ingresen exactamente 2 cadenas
		if len(input) != 2 {
			fmt.Println("Error: Debes ingresar exactamente el nombre del animal y la acción.")
			continue
		}

		name, action := input[0], input[1]
		animal, exists := animals[name]

		if !exists {
			fmt.Println("Error: Animal no encontrado (vaca, pájaro, serpiente).")
			continue
		}

		// Ejecutamos el método correspondiente
		switch action {
		case "comer":
			animal.Eat()
		case "moverse":
			animal.Move()
		case "hablar":
			animal.Speak()
		default:
			fmt.Println("Error: Acción no válida (comer, moverse, hablar).")
		}
	}
}