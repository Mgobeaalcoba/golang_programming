// Importante comenzar con el package main ya que indica por donde arranca la ejecución de mi paquete.
package main

import "fmt"

func main() {
	// go es un lenguaje tipado: se especifica el tipo de dato
	// 1° forma: declarar y luego inicializar la variable
	var nombre1 string
	nombre1 = "Mariano"

	// 2° forma: declarar e inicializar la variable a la vez
	var nombre2 string = "Nicole"

	// 3° forma: declarar e inicializar a la vez de forma corta
	edad := 35

	fmt.Println("Mi nombre es", nombre1, "mi mujer se llama", nombre2, "y mi edad es de", edad, "años. Hola Mundo!!!")

	var a int
	var b float64
	var c string
	var d bool

	// En golang las variables ya vienen con valores por defecto. No vienen undefined como en Kotlin por ejemplo.

	fmt.Println(a,b,c,d) // Imprime: 0 0  false

	const pi = 3.141592 // Las constantes no pueden ser modificadas posteriormente. 

	fmt.Println(pi)
}