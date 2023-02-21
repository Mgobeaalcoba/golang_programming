// Funcion recursiva: Es una función que se ejecuta a si misma.

package main

import "fmt"

func factorial(n int) int {
	// factorial de 0 es 1
	if n == 0 {
		return 1
	}

	f := n * factorial(n-1)

	return f
}

func main() {
	// Función recibe un numero, por ej un 3 y debe multiplicarse ese numero por todos los anteriores hasta llegar al 1. En este caso sería func(3) = 3*2 + 3*1
	fmt.Println(factorial(15))
}
