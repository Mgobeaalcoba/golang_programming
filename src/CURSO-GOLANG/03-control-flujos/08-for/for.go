package main

import "fmt"

// En golang no existe el while. Hay que generarlo desde el for.

func main() {
	// Bucle infinito
		/*
		for {
			fmt.Println("Hola Mundo")
			// Imprime al infinitum "Hola Mundo". Si no lo corto con Ctrl + C va a terminar con mi memoria RAM y tildar la PC
		}
		*/

	// Bucle tipo while
		numeros := 1245556549865
		numeros2 := numeros
		c := 0 
		for numeros > 0 {
			numeros /= 10
			c++ // Me va a dar la cantidad de caracteres que tiene el numero "numeros"
		}
		fmt.Println(numeros2, "tiene", c, "caracteres")

	// Bucle for
		for i := 0; i < 10; i++ {
			if i % 2 == 0 {
				fmt.Println("Hola", i)
			} 
		}
}