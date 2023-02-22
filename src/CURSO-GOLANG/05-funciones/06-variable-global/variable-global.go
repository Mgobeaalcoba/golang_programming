// Las variables globales están por fuera de cualquier función, incluida la función main(). Si son modificadas, entonces la modificación va a hacer efecto en todas las funciones en las que dicha variable sea invocada.

// Tmb se pueden crear const globales, pero las mismas no se pueden modificar.

package main

import "fmt"

// Variable global
var mensaje string

// Constante global
const mensaje2 string = "Este msj no se puede modificar!!!"

func function1() {
	mensaje = "Hola desde la función uno!"
	fmt.Println(mensaje)
}

func function2() {
	mensaje = "Hola desde la función dos!"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la función principal"
	fmt.Println(mensaje)
	fmt.Println(mensaje2)

	defer function1() // Si uso defer antes de invocar una función la misma se va a ejecutar al final de la función.
	// defer es util para asincronismo por ejemplo.
	function2()

	fmt.Println(mensaje)
}

/*
Impresión sin defer:

Hola desde la función principal
Este msj no se puede modificar!!!
Hola desde la función uno!
Hola desde la función dos!
Hola desde la función dos!

Impresión con defer:

Hola desde la función principal
Este msj no se puede modificar!!!
Hola desde la función dos!
Hola desde la función dos!
Hola desde la función uno!
*/
