package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		// Continue: Salta la iteración en la que estás:
		if i == 5 {
			fmt.Println("Se salta la iteración 5")
			continue
		}

		// Break: Corta el ciclo
		if i == 8 {
			fmt.Println("Romper con el bucle")
			break
		}
		fmt.Println(i)

	}
}
