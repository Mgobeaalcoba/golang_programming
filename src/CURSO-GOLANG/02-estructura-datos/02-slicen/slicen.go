package main

import "fmt"

func main() {

	// Slicen: a diferencia de los arrays son mutables. Podemos agregarles nuevos elementos o quitarles elementos
	// Para definirlos se deja el espacio entre corchetes vacio
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	// Funcion append para agregar elementos a un slicen

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	// Sub slicen

	subSlicen := numeros[0:2]

	fmt.Println(subSlicen)

	numeros[0] = 100

	fmt.Println(subSlicen) // Al mutar numeros muta tmb subSlicen
	fmt.Println(numeros)

	// Caracteristicas de los slicen: 
	// Punteros: Donde inicia y donde termina
	// Longitud: Len o largo de un slicen
	// Capacidad:

	meses := []string{"Enero","Febrero","Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
	// Len: 3, Cap: 3, 0xc0000221b0

	meses = append(meses, "Abril")

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
	// Len: 4, Cap: 6, 0xc00005e060

	// Se alarga el len en +1, pero se duplica la capacidad de 3 a 6. Mientras que la referencia en memoría pasa a ser otra distinta. Estoy creando otra slicen en memoria

}