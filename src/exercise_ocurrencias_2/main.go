package main

import "fmt"

func main() {
	str1 := "holahola"
	mapa := make(map[rune]int) // Estructura map vacia. Es como un diccionario para Python. Clave rune = int32 y valor int

	/*
		Luego, se utiliza un bucle for para iterar sobre cada carácter en la cadena de texto str1. Dentro del bucle, se utiliza un condicional if para verificar si el carácter actual ya existe en el mapa mapa. Si el carácter aún no se ha encontrado en el mapa, se agrega como una nueva clave en el mapa y se establece su valor en 1. Si el carácter ya existe en el mapa, se incrementa su valor en 1.
	*/
	for _, v := range str1 {
		if _, ok := mapa[v]; !ok {
			mapa[v] = 1
		} else {
			mapa[v] += 1
		}
	}

	// Uso el metodo string para transformar la letra en unicode (numero) en letra e imprimirlo transformado
	for k, v := range mapa {
		fmt.Printf("%s:%d ", string(k), v)
	}
}