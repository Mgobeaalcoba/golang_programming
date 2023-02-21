// ¿Que es un closure? Es una función que retorna a otra función, y dentro de esa función podemos recordar el valor de la función padre.

package main

import (
	"fmt"
	"strings"
)

// Función closure que va a recibir un entero y va a retornar otra función que a su vez va a recibir un string y retornar un string
// Objetivo del closure: Que el usuario ingrese un entero, y eso determine la cantidad de veces que se va a repetir una cadena de texto por consola.
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	// Función anónima:
	func() {
		fmt.Println("Hola desde la función anónima")
	}() // Los parentesis que abren y cierran al final autoejecutan a esta función

	myFunc := func(nombre string) string {
		saludo := fmt.Sprintf("Hola %s, desde la segunda función anonima", nombre)
		return saludo
	}
	// Para invocarla debo llamarla como a una función normal
	fmt.Println(myFunc("Mariano"))
	// Si solo llamo myFunc sin parentesis me muestra el espacio en memoria que ocupa
	//fmt.Println(myFunc) // 0x4805c0

	// Función Closure:

	repeat3 := repeat(3) // Guardo en repeat3 la cantidad de veces que quiero repetri mi palabra. repeat 3 es una función de repeat con n = 3. Luego tengo que reutilizarla abajo para que repita la palabra que quiero
	fmt.Println(repeat3("Hola "))

	// Si en lugar de repetir 3 veces necesito repetir 10, entonces guardo mi closure en una nueva función a la que voy a llamar repeat10

	repeat10 := repeat(10)
	fmt.Println(repeat10("Suerte "))

}
