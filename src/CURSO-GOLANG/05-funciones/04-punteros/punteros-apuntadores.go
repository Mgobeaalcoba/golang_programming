// Punteros: Claves
// 1- Al parametro de la función debo agregarle un * antes del tipo de dato
// 2- Al argumento que envio al invocar la función debo agregarle un & antes del nombre de la función.

package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena) // Si trabajo con punteros no necesito usar el &
	*cadena = "Hola desde la función"
}

func main() {
	cadena := "Hola Mundo de Go"
	fmt.Printf("%p\n", &cadena) // El & se usa para imprimir la referencia de la memoria
	fmt.Println("Antes de la función:", cadena)

	//modificarValor(cadena) //Estamos pasando una copia de cadena

	//fmt.Println("Despues de la función:", cadena)

	// Si quiero pasar la variable en memoria tal cual entonces debo usar punteros:

	modificarValor(&cadena) // Paso la variable original con puntero

	fmt.Println("Despues de la función:", cadena)

	// Modificamos el valor de la variable "cadena" mediante una función "modificarValor"

}
