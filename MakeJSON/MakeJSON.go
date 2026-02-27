package main
import (
	"encoding/json"
	"fmt"
)
// Definimos la estructura del objeto que queremos crear
type Persona struct {
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}
func main() {
	var n, d string
	// 1. Leemos el nombre y la dirección
	// El sistema suele introducir los datos uno tras otro
	fmt.Scan(&n)
	fmt.Scan(&d)
	// 2. Creamos una instancia de nuestra estructura con los datos recibidos
	p := Persona{Nombre: n, Direccion: d}
	// 3. Convertimos la estructura a formato JSON (Marshalling)
	datosJSON, err := json.Marshal(p)
	if err != nil {
		return
	}
	// 4. Imprimimos el objeto JSON como una cadena de texto
	fmt.Println(string(datosJSON))
}