/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "") 
	arrayInvertida := make([]string, 0) 

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	cadenaInvertida := strings.Join(arrayInvertida, "")

	return cadenaInvertida
}

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")

	if palabra == reverse(palabra) {
		return true
	} else {
		return false
	}

}

func main() {

	fmt.Println(esPalindromo("Oso"))

	fmt.Println(esPalindromo("Luz azul"))

	fmt.Println(esPalindromo("Mariano"))

	fmt.Println(esPalindromo("Neuquen"))

	fmt.Println(esPalindromo("Nicole"))

}