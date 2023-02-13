package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Print(hola )
	fmt.Print(mundo)
	// Print imprime sin dejar una linea entre impresión e impresión

	fmt.Println(hola)
	fmt.Println(mundo)
	// Print line imprime y luego da un salto de linea

	nombre := "Mariano"
	edad := 35

	fmt.Printf("Hola %s y su edad es %d \n", nombre, edad)
	// %s se usa para variables string - %d se usa para variables enteras
	fmt.Printf("Hola %v y su edad es %v \n", nombre, edad)
	// %v se usa para variables que desconocemos o no queremos declarar el tipo de dato.
	// Printf imprime con variables al interior de la linea

	mensaje := fmt.Sprintf("Hola %s y su edad es %d", nombre, edad)
	// Sprintf sirve para guardar una linea de texto formateada y con variables dentro de otra variable
	fmt.Println(mensaje)

	// Con Printf tmb puedo saber de que tipo es una variable. Por ej:
	fmt.Printf("nombre: %T \n", nombre) // Pregunto por el tipo de variable que es nombre. 

	// Con fmt tmb puedo recibir valores por consola, el input en python
	fmt.Print("Ingres otro nombre: ")
	fmt.Scanln(&nombre) // Capturo en la variable nombre, ya inicializada lo que recibo por consola.

	println()
	
	println("Otro nombre", nombre)
}