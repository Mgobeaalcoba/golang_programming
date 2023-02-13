package main

import "fmt"

func main() {

	// Arrays vacios

	// arrays: en golang pueden almacenar un único tipo de datos y debe predefinirse por ser un lenguaje de tipado estatico - Tmb debo definir la cantidad de elementos de mi arreglo entre []

	var numeros [5]int
	fmt.Println(numeros)

	// Si no pongo nada en mi array al imprimirse sale [0 0 0 0 0]

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)
	fmt.Println(numeros[3])

	// Arrays con valores

	nombres := [4]string{"Mariano", "Nicole", "Lisandro", "Lautaro"}
	fmt.Println(nombres)

	// Arrays sin definir su longitud

	colores := [...]string{"Negro", "Rojo", "Verde", "Amarillo"}
	fmt.Println(colores)
	fmt.Println("La longitud de mi arreglo es de", len(colores))

	// Arrays con indices indefinidos

	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println(monedas, len(monedas))

	// El largo será 6 dado que dejamos sin definir el valor 1 y 4 del array

	monedas[1] = "Rublos"

	fmt.Println(monedas, len(monedas))

	monedas[4] = "Libras"

	fmt.Println(monedas, len(monedas))

	// Sub-Arrays o Slicing:

	submonedas := monedas[0:3] // El último valor queda excluido como en python

	fmt.Println(submonedas)

	// Si quiero hacer un sub array hasta el final entonces dejo el segundo valor vacio

	submonedas2 := monedas[3:]

	fmt.Println(submonedas2)

}
