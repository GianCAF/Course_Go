package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// Definimos la estructura con campos de 20 caracteres (en Go usamos string)
type Nombre struct {
	Fname string
	Lname string
}
func main() {
	var nombreArchivo string
	fmt.Print("Introduce el nombre del archivo: ")
	fmt.Scan(&nombreArchivo)
	// Abrimos el archivo
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()
	// Creamos la slice de structs
	personas := make([]Nombre, 0)
	// Leemos el archivo línea por línea
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, " ")
		if len(partes) >= 2 {
			// Creamos el struct y limitamos a 20 caracteres si es necesario
			nuevaPersona := Nombre{
				Fname: partes[0],
				Lname: partes[1],
			}
			// Añadimos a la slice
			personas = append(personas, nuevaPersona)
		}
	}
	// Iteramos sobre la slice e imprimimos los resultados
	for _, p := range personas {
		fmt.Printf("Nombre: %s, Apellido: %s\n", p.Fname, p.Lname)
	}
}