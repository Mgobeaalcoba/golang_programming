/*
Funciones variaticas:
Son aquellas que reciben argumentos indefinidos.
*/

package main

import "fmt"

// Parametros indeterminados se marcan con ...type
// Parametro determinados solo con el nombre y tipo
// Si los combino entonces debo poner primero los determinados y luego los indeterminados

// Para returnar mas de un tipo de valor debo especificarlo entre parentesis.

func sumar(nombre string, numeros ...int) (string, int) {
	// Al poner ... los numeros que reciba los guardo en un slicen
	fmt.Println(numeros)
	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	/* Mi solución:
	var suma int
	for i := 0; i < len(numeros); i++ {
		suma += numeros[i]
	}
	fmt.Printf("La suma es %d \n", suma)
	*/
	var total int
	for _, num := range numeros {
		total += num
	}

	return mensaje, total

}

func main() {
	mensaje, result := sumar("Mariano", 5, 10, 15, 20)
	fmt.Printf("%s %d \n", mensaje, result)

	mensaje, result = sumar("Nicole", 10, 15, 20, 33, 45, 75)
	fmt.Printf("%s %d \n", mensaje, result)
}
