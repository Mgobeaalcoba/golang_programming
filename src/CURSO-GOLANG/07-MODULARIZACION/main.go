package main

import (
	"fmt"
	"paquetes/mensajes"
)

func main() {
	mensajes.Hola()
	mensajes.Imprimir()

	// Pero no puedo acceder a mi variable "mensaje"

	//fmt.Println(mensajes.mensaje) // mensaje not exported by package mensajescompilerUnexportedName

	// Si la declaro con mayuscula si puedo. Ejemplo: 

	fmt.Println(mensajes.Mensaje2)

	// mensajes.funcionPrivada() // mensaje not exported by package mensajescompilerUnexportedName

	// Tengo que correr la función privada desde el propio archivo saludar.go u otro archivo del paquete mensajes para que se use esa función. 
}

// Si mis carpetas no tuvieran guiones, entonces podría solo poniendo en mi func main "saludar." importar el paquete de saludar. Pero como tienen guiones entonces debo importarlos manualmente.

// En este caso al ejecutar main.go voy a estar llamando y usando tambien mi archivo saludar.go de mi paquete mensajes al correrlo. 

// Recordar que en la raiz del proyecto. En este caso "07-MODULARIZACION" debo correr por terminal previamente: 

// go mod init paquetes --> El nombre final puede variar. 

// Output:

// C:\Users\mgobea\go\src\CURSO-GOLANG\07-MODULARIZACION(main -> origin)
// λ go run main.go
// Hola desde el paquete mensaje