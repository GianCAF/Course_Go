package main
import (
	"fmt"
	"strings"
)
func main() {
	var texto string
	// El sistema ingresará la cadena aquí
	fmt.Scan(&texto)
	// Convertimos a minúsculas para asegurar que encuentre 'i', 'a', 'n'
	texto = strings.ToLower(texto)
	// Verificamos la presencia de las tres letras
	if strings.Contains(texto, "i") && strings.Contains(texto, "a") && strings.Contains(texto, "n") {
		fmt.Println("¡Encontrado!")
	} else {
		fmt.Println("¡No encontrado!")
	}
}