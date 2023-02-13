package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool // Guarda el resultado de la comparación

	// Igualdad
	r = a == b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)
	// 4 es igual que 5? false

	a = 3
	b = 3

	r = a == b

	fmt.Printf("%d es igual que %d? %t \n", a, b, r)
	// 3 es igual que 3? true

	// Distinto

	r = a != b // Se escribe ! = todo junto
	fmt.Printf("%d es distinto que %d? %t \n", a, b, r)
	// 3 es distinto que 3? false

	// Mayor que

	r = a > b // Se escribe ! = todo junto
	fmt.Printf("%d es mayor que %d? %t \n", a, b, r)

	// Menor que

	r = a < b // Se escribe ! = todo junto
	fmt.Printf("%d es menor que %d? %t \n", a, b, r)

	// Mayor o igual que

	r = a >= b // Se escribe ! = todo junto
	fmt.Printf("%d es mayor o igual que %d? %t \n", a, b, r)

	// Menor o igual que

	r = a <= b // Se escribe ! = todo junto
	fmt.Printf("%d es menor o igual que %d? %t \n", a, b, r)

}