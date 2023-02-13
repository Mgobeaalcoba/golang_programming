package main

import "fmt"

func main() {
	fmt.Println("Ingrese el primer numero entero a sumar")
	var n1 int
	fmt.Scanln(&n1)
	fmt.Println("Ingrese el segundo numero entero a sumar")
	var n2 int
	fmt.Scanln(&n2)

	var total int = suma(n1, n2)
	fmt.Println("La suma es", total)
}

func suma(num1, num2 int) int {
	return num1 + num2
}
