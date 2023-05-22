package main

import (
	"fmt"
	"math/rand"
)

func esPar(numero int) bool {
	if numero%2 == 0 {
		return true
	}
	return false
}

func main() {
	// Ejemplo de uso de funciones
	//genera un numero aleatorio
	//generate a random number
	numero := rand.Intn(100)

	if esPar(numero) {
		fmt.Println(numero, "es par")
	} else {
		fmt.Println(numero, "es impar")
	}
}
