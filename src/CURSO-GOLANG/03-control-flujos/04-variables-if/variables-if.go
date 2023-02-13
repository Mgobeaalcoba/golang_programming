// En golang, a diferencia de otros lenguajes es posible crear variables cuyo scope o alcance es solamente el condicional en el que viven.

package main

import "fmt"

func main() {

	// Abajo, inicializo y declaro las variables en el mismo condicional if
	if nombre, edad := "Alex", 26; nombre == "Alex" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d", nombre, edad)
	}

	// println(nombre,edad) // undefined: nombre
	// Si deseo invocarlas por fuera del scope del if tendré un error con el de arriba

	if nombre, edad := "Mariano", 35; nombre == "Alex" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d \n", nombre, edad)
	}

	// Obtener valor de mapa
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	fmt.Println(users["Mariano"]) // Devuelve ""
	fmt.Println(users["Alex"]) // Devuelve alex@gmail.com

	// Con variables puedo guardar la devolución que me hace golang para cada key. 

	correo, ok := users["Juan"]
	fmt.Println(correo, ok) // en correo guarda el contenido si existe. En ok guarda true si existe y false si no existe
	//  false

	correo, ok = users["Alex"]
	fmt.Println(correo, ok)
	// alex@gmail.com true

	// Esto podemos usarlo para validar que una key exista antes de determinar un camino a seguir

	if correo, ok = users["Juan"]; ok == true {
		// "ok == true" se puede poner solo como "ok"
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener su valor.")
	} // No fue posible obtener su valor

	if correo, ok = users["Alex"]; ok == true {
		// "ok == true" se puede poner solo como "ok"
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener su valor.")
	} // alex@gmail.com

}