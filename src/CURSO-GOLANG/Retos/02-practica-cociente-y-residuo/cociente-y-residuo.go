package main

import "fmt"

func main() {
	fmt.Println("Ingrese el primer numero entero a dividir")
	var n1 int
	fmt.Scanln(&n1)
	fmt.Println("Ingrese el segundo numero entero a dividir")
	var n2 int
	fmt.Scanln(&n2)

	var total int = dividir(n1, n2)
	fmt.Println("El cociente es", total)

	var total2 int = residuo(n1, n2)
	fmt.Println("El residuo es", total2)
}

func dividir(num1, num2 int) int {
	return num1 / num2
}

func residuo(num1, num2 int) int {
	return num1 % num2
}
