package main
import "fmt"
func main() {
    var numero float64
    // 1. Leemos el número decimal
    fmt.Scan(&numero)
    // 2. Lo convertimos a int (esto trunca los decimales)
    resultado := int(numero)
    // 3. Imprimimos el resultado
    fmt.Println(resultado)
}