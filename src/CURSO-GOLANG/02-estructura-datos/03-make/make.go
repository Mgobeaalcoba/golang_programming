package main

import "fmt"

func main() {

	numeros := make([]int, 3, 3)
	// Parametros de make: tipo de dato, largo y capacidad

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	// Sino mantengo la longitud de 0 me saldrá error. Porque no existen esos indices en mi slicen. Debo cambiar la longitud a 3 para poder hacerlos
	// panic: runtime error: index out of range [0] with length 0

	fmt.Println(numeros, len(numeros), cap(numeros))
	// [100 200 300] 3 3

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))
	// [100 200 300 400] 4 6
	
}