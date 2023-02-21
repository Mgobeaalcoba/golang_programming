package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	// En go no hay función para revertir. Hay que hacer varios pasos mas.
	// Split convierte una cadena en arreglo
	//arrayCadena := strings.Split(cadena, " ") // Splitea por espacio en blanco
	arrayCadena := strings.Split(cadena, "") // Splitea por cada espacio en el string
	arrayInvertida := make([]string, 0)      // Armo un slicing vacio para luego almacenar el invertido

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	//fmt.Println(arrayCadena)
	//fmt.Println(arrayInvertida)

	// Con join guardo en otra variable el resultado de la inversión de mi cadena/array para luego retornarla
	cadenaInvertida := strings.Join(arrayInvertida, "")
	// El join se hace con el conector declarado como segundo parametro.

	//fmt.Println(cadenaInvertida)

	return cadenaInvertida
}

func esPalindromo(palabra string) bool {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	//fmt.Println(palabra)
	//palabra = strings.Replace(palabra, " ", "", 1)
	palabra = strings.ReplaceAll(palabra, " ", "")
	//fmt.Println(palabra)

	if palabra == reverse(palabra) {
		fmt.Println("Es un palindromo")
		return true
	} else {
		fmt.Println("No es un palindromo")
		return false
	}

}

func main() {
	// Palindromo es aquella palabra que se lee igual comenzando desde la derecha o desde la izquierda. Ejemplo: "Neuquen"
	esPalindromo("Oso")

	// "Luz azul" tmb es un palindromo si eliminamos el espacio en el medio. ¿Como?
	esPalindromo("Luz azul")

	esPalindromo("Mariano")

	esPalindromo("Neuquen")

	esPalindromo("Nicole")

}
