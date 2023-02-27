/*
Ocurrencias de cada carácter en un string (sin tener en cuenta caracteres especiales, solo letras).
Si tenemos un string: “bbacdcbba”
Queremos como respuesta (no interesa el orden):
[]Occurrence{
    {Letter: ‘a’, Times: 2},
    {Letter: ‘b’, Times: 4},
    {Letter: ‘c’, Times: 2},
    {Letter: ‘d’, Times: 1},
}
*/

package main

import (
	"fmt"
	"strings"
)

func contador(letra string, cadena []string) int{
	repeat := 0
	for _, elemento := range cadena{
		if letra == elemento {
			repeat += 1
		}
	}
	return repeat
}

func main() {
	fmt.Println("Ingrese una cadena o string: ")
	var str1 string
	fmt.Scanln(&str1)
	arrayCadena := strings.Split(str1, "")
	arrayVacia := []string{}
	repeat := false
	for _, indice := range arrayCadena {
		for _,indice2 := range arrayVacia {
			if indice == indice2 {
				repeat = true
			}
		}
		if repeat == false {
			cant := contador(indice, arrayCadena)
			fmt.Println("Letter:",indice,"Times:",cant )
		}
		repeat = false
		arrayVacia = append(arrayVacia, indice)
	}
}