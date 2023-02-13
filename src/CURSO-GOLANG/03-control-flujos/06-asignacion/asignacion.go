package main

import "fmt"

func main() {

	a := 2

	a = a + 2 // Reasignación sin operadores

	fmt.Println(a) // 4

	// Operadores de asignación

	// += -= *= /= %=

	// Suma en asignación

	a += 2 // Lo mismo que a = a + 2

	fmt.Println(a) // 6

	// Resta en asignación

	a -= 2 // Lo mismo que a = a + 2

	fmt.Println(a) // 4

	// Multiplicación en asignación

	a *= 2 // Lo mismo que a = a + 2

	fmt.Println(a) // 8

	// División en asignación

	a /= 2 // Lo mismo que a = a + 2

	fmt.Println(a) // 4

	// Modulo de asignación

	a %= 2 // Lo mismo que a = a + 2

	fmt.Println(a) // 0

}