package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)

	switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "es vocal")
	case vocal == "e" || vocal == "E":
		fmt.Println(vocal, "es vocal")
	case vocal == "i" || vocal == "I":
		fmt.Println(vocal, "es vocal")
	case vocal == "o" || vocal == "O":
		fmt.Println(vocal, "es vocal")
	case vocal == "u" || vocal == "U":
		fmt.Println(vocal, "es vocal")
	default:
		fmt.Println(vocal, "no es vocal")
	}

	// Otra forma de trabajar con switch, mas practica para el caso de uso de arriba es:

	switch vocal {
	case "a", "e", "i", "o", "u":
		fmt.Println(vocal, "es vocal minuscula")
	case "A", "E", "I", "O", "U":
		fmt.Println(vocal, "es vocal MAYUSCULA")
	default: 
		fmt.Println(vocal, "no es vocal")
	}
}