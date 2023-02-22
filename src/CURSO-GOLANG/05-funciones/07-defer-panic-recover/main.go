package main

import (
	"fmt"
	//"log"
	"os"
)

func main() {

	// Recover: Sirve para que nuestra aplicación no se detenga de forma brusca frente a un panico.
	defer func() {
		// recover me va a devolver si hubo algún panic en mi aplicación. error != nil significa que ocurrio un panic:
		if error := recover(); error != nil {
			fmt.Println("Ups! al parecer el programa no finalizo de forma correcta!!!")
			// El recover reemplaza el contenido del panic y nos permite continuar con nuestro codigo cuando en la ejecución del mismo hubo algún panic.
		}
	}()

	/*
		El codigo de abajo se puede optimizar y agregar en el if:

		file, error := os.Open("hola.txt") // Nos devuelve una variable de tipo file y una de tipo error.

	*/

	if file, error := os.Open("src/CURSO-GOLANG/05-funciones/07-defer-panic-recover/hola.txt"); error != nil {

		// log.Fatal(error)
		// fmt.Println(error)
		// fmt.Println("No fue posible obtener el archivo!!")
		// Puedo llamar ese msj de arriba desde panic que luego me cierra el programa así:
		panic("No fue posible obtener el archivo!!")

	} else {

		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}() // Se va a ejecutar al final del else

		// como leo el archivo
		contenido := make([]byte, 254)               // Creo un slicen de tipo byte
		long, _ := file.Read(contenido)              // Leo el archivo y guardo su long y errores en dos variables
		contenidoArchivo := string(contenido[:long]) // Convierto en un string el contenido del archivo que tenía en bytes desde el inicio hasta el largo del archivo que obtuve arriba.
		fmt.Println(contenidoArchivo)

	}

	// Si abrí una conexión con un archivo es necesario cerrarla al finalizar el programa para no dejar RAM usada sin necesidad. (optimización)

	//file.Close() // La tengo que llevar dentro del if/else para que funcione. Pero la voy a llevar con "defer"
}

// Por algún motivo que desconozco no logro hacer correr os en WSL2(Ubuntu). Estoy averiguando el motivo y vuelvo.

// El problema de arriba está resuelto. Ya que al abrir code desde la carpeta padre "go" entiende que "hola.txt" debería estar en esa raiz. Por lo que para llegar debo guiarel contenido de "os.Open(./...)" desde go hasta donde se encuentra.
