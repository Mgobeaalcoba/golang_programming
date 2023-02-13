package main

import "fmt"

// Operadores que nos permiten incrementar y decrementar en 1 a un entero.

func main() {

	a := 0
	fmt.Println(a) // 0

	// Operador de incremento
	a++
	fmt.Println(a) // 1

	// Operador de decremento
	a--
	fmt.Println(a) // 0

	// Incremento en 5
	a++
	a++
	a++
	a++
	a++
	fmt.Println(a) // 5

}