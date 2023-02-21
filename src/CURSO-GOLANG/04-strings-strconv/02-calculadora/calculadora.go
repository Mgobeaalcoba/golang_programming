package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	// 4+4+5+8
	valores := strings.Split(expresion, "+")
	fmt.Printf("%T \n", valores) // Devuelve el tipo de datos de mi array
	fmt.Println(valores)

	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor) // Atoi de strconv transforma de string a entero.
		// result = result + num
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("Solo debes utilizar el operador +")
		} else {
			result += num
			//fmt.Println(num, error)
		}
		// Si error devuelve <nil> entonces estamos bien. Sino marca un error.
		// strconv.Atoi: parsing "4*5": invalid syntax
	}

	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=> ")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("=> %d \n", result)

	// En Golang no existe el manejo de errores por lo que lo tenemos que hacer a mano revisando el codigo.

}
