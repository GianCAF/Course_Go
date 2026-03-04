package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal es la interfaz que describe las capacidades de un animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Estructuras de datos para cada tipo de animal
type Cow struct{}
func (c Cow) Eat()   { fmt.Println("hierba") }
func (c Cow) Move()  { fmt.Println("caminar") }
func (c Cow) Speak() { fmt.Println("mugido") }

type Bird struct{}
func (b Bird) Eat()   { fmt.Println("gusanos") }
func (b Bird) Move()  { fmt.Println("volar") }
func (b Bird) Speak() { fmt.Println("pío") }

type Snake struct{}
func (s Snake) Eat()   { fmt.Println("ratones") }
func (s Snake) Move()  { fmt.Println("deslizarse") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	// Mapa para almacenar los animales creados por el usuario
	animalMap := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bienvenido al Gestor de Animales")
	fmt.Println("Comandos: 'newanimal <nombre> <tipo>' o 'query <nombre> <accion>'")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.Fields(strings.ToLower(scanner.Text()))
		if len(input) != 3 {
			fmt.Println("Error: Se requieren exactamente 3 parámetros.")
			continue
		}

		command, name, param := input[0], input[1], input[2]

		switch command {
		case "newanimal":
			// Lógica para crear un nuevo animal según su tipo
			var newAni Animal
			switch param {
			case "vaca":
				newAni = Cow{}
			case "pájaro":
				newAni = Bird{}
			case "serpiente":
				newAni = Snake{}
			default:
				fmt.Println("Error: Tipo de animal inválido (vaca, pájaro, serpiente).")
				continue
			}
			animalMap[name] = newAni
			fmt.Println("¡Creado!")

		case "query":
			// Lógica para consultar un animal existente
			ani, exists := animalMap[name]
			if !exists {
				fmt.Println("Error: El animal con ese nombre no existe.")
				continue
			}

			switch param {
			case "comer":
				ani.Eat()
			case "moverse":
				ani.Move()
			case "hablar":
				ani.Speak()
			default:
				fmt.Println("Error: Acción inválida (comer, moverse, hablar).")
			}

		default:
			fmt.Println("Error: Comando no reconocido (newanimal o query).")
		}
	}
}