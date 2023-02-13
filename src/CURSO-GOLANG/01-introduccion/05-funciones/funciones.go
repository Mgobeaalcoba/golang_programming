package main

import "fmt"

// funciones basicas: que se ejecuten y que retornen

func saludar(nombre string) {
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int {
	return a + b
}

// En los parentesis van las variables que serán inputs/argumentos y su tipo. Entre los parentesis y la llave debo declarar el tipo de dato que retorna mi función en golang

func main() {
	saludar("Mariano")
	r := sumar(10, 20)
	fmt.Println(sumar(5, 7))
	fmt.Println(r)
}
