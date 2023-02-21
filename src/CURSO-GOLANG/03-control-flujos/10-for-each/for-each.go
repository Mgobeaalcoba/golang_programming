package main

import "fmt"

func main() {
	nombres := [...]string{"Alex", "Roel", "Juan"}

	// for each itera los datos de un arreglo

	// con un for normal lo hariamos así:

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	// con un for each: Declaro el arreglo a iterar luego del range

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	// Si solo quieero el elemento reemplazo el indice por un "_"
	for _, elemento := range nombres {
		fmt.Println(elemento)
	}

	// Si solo quieero el indice no es necesario el "_" dado que el primer valor del for each es siempre el indice
	for indice := range nombres {
		fmt.Println(indice)
	}
}
