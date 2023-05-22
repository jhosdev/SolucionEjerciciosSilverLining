package main

import "fmt"

func esPar(numero int) bool {
    if numero%2 == 0 {
        return true
    }
    return false
}

func main() {
    // Ejemplo de uso de funciones
    //genera un numero aleatorio
    numero := rand.Intn(100)
    
    if esPar(numero) {
        fmt.Println(numero, "es par")
    } else {
        fmt.Println(numero, "es impar")
    }
}
