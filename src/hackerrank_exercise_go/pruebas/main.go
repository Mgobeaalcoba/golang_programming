package main

import (
	"fmt"
	"strings"
)

func biggerIsGreater(w string) string {

	arrayString := strings.Split(w, "")
	copiaArray := make([]string, len(arrayString)) // Armo un array vacio independiente del anterior
	copy(copiaArray, arrayString)                  // Copio mi array original en el segundo para luego compararlos
	var init int = len(arrayString)

	for i := len(arrayString) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			count1j := 1
			if arrayString[i] > arrayString[j] {
				aux := arrayString[j]
				arrayString[j] = arrayString[i]
				arrayString[i] = aux
				init = j + 1
				fmt.Println("J vale:", j, ".I vale:", i)
				fmt.Println("Entro en este primer ciclo:", count1j, "veces")
				fmt.Println("Así está el string recibido en el primer ciclo:", arrayString)
				count1j += 1
				break
			}
		}
		fmt.Println("copiaArray:", copiaArray)
		fmt.Println("arrayString:", arrayString)
		if strings.Join(copiaArray, "") != strings.Join(arrayString, "") {
			break
		}
	}
	if init < len(arrayString)-1 {
		for i := init; i < len(arrayString); i++ {
			//count2i := 1
			for j := i + 1; j < len(arrayString); j++ {
				count1j := 1
				if arrayString[i] > arrayString[j] {
					aux := arrayString[j]
					arrayString[j] = arrayString[i]
					arrayString[i] = aux
					fmt.Println("Entro en este segundo ciclo:", count1j, "veces")
					fmt.Println("Así está el string recibido:", arrayString)
					count1j += 1
				}
			}
		}
	}
	if strings.Join(copiaArray, "") != strings.Join(arrayString, "") {
		return strings.Join(arrayString, "")
	} else {
		return "no answer"
	}
}

func main() {
	cadena := "dkhc" // Respuesta esperada: hcdk // No dkch
	fmt.Println("Cadena enviada:", cadena)
	fmt.Println("Cadena devuelta:", biggerIsGreater(cadena))
	fmt.Println()
	cadena2 := "dhck" // Respuesta esperada: dhkc // No dkch
	fmt.Println("Cadena2 enviada:", cadena2)
	fmt.Println("Cadena2 devuelta:", biggerIsGreater(cadena2))
	fmt.Println()
	cadena3 := "hefg" // Respuesta esperada: hegf // No dkch
	fmt.Println("cadena3 enviada:", cadena3)
	fmt.Println("cadena3 devuelta:", biggerIsGreater(cadena3))
	fmt.Println()
	cadena4 := "bb" // Respuesta esperada: bb o "no answer" // No dkch
	fmt.Println("cadena4 enviada:", cadena4)
	fmt.Println("cadena4 devuelta:", biggerIsGreater(cadena4))
	fmt.Println()

	cadena5 := "abdc" // Respuesta esperada: acbd // No acdb
	fmt.Println("cadena5 enviada:", cadena5)
	fmt.Println("cadena5 devuelta:", biggerIsGreater(cadena5))
	fmt.Println()

	cadena6 := "xildrrhpca" // Respuesta esperada: xildrrpach
	fmt.Println("cadena6 enviada:", cadena6)
	fmt.Println("cadena6 devuelta:", biggerIsGreater(cadena6))
	fmt.Println()

}
