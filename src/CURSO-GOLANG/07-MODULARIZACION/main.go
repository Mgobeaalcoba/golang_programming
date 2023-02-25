package main

import (
	"fmt"
	"paquetes/models" // Importo un paquete propio luego de activarlo con go mod init ...

	figuras "github.com/Mgobeaalcoba/geometrics_golang" // Importo mi propio modulo subido a github
	"github.com/donvito/hellomod"                       // Importo un paquete de terceros luego de instalarlo con go get...
)

func main() {
	/*
		mensajes.Hola()
		mensajes.Imprimir()
	*/

	p1 := models.Persona{} // No puedo cargar los atributos por ser privados. Voy a tener que crear un constructor junto con la struct persona.

	// Como los atributos son privados, debo contruir mi objeto con su constructor:
	p1.Constructor("Mariano", 35)

	fmt.Println(p1)

	// Si quiero modificar un valor, no puedo hacerlo como antes:

	// p1.nombre = "Daniel" // p1.nombre undefined (type models.Persona has no field or method nombre)

	// Sino que debo volver a usar mi constructor:

	p1.Constructor("Daniel", 35)

	fmt.Println(p1)

	// Si por el contrario, deseo solo modificar uno solo debo crear un metodo de encapsulación o encapsulamiento... Getter (Get) y Setter (Set) se conocen en otros lenguajes de programación

	fmt.Println(p1.GetNombre())
	fmt.Println(p1.GetEdad())

	p1.SetEdad(18)
	fmt.Println(p1)
	p1.SetNombre("Francisco")
	fmt.Println(p1)

	/*
		// Pero no puedo acceder a mi variable "mensaje"

		//fmt.Println(mensajes.mensaje) // mensaje not exported by package mensajescompilerUnexportedName

		// Si la declaro con mayuscula si puedo. Ejemplo:
	*/

	/*
		fmt.Println(mensajes.Mensaje2)
	*/

	/*
		// mensajes.funcionPrivada() // mensaje not exported by package mensajescompilerUnexportedName

		// Tengo que correr la función privada desde el propio archivo saludar.go u otro archivo del paquete mensajes para que se use esa función.
	*/

	hellomod.SayHello() // Uso una función del paquete descargado.

	cir1 := figuras.Circulo{Radio: 10.2, Pi: 3.14}

	fmt.Println(cir1)

	fmt.Print("El area de mi circulo es: ")
	fmt.Println(figuras.CalculoDeArea(&cir1))

}

/*
// Si mis carpetas no tuvieran guiones, entonces podría solo poniendo en mi func main "saludar." importar el paquete de saludar. Pero como tienen guiones entonces debo importarlos manualmente.

// En este caso al ejecutar main.go voy a estar llamando y usando tambien mi archivo saludar.go de mi paquete mensajes al correrlo.

// Recordar que en la raiz del proyecto. En este caso "07-MODULARIZACION" debo correr por terminal previamente:

// go mod init paquetes --> El nombre final puede variar.

// Output:

// C:\Users\mgobea\go\src\CURSO-GOLANG\07-MODULARIZACION(main -> origin)
// λ go run main.go
// Hola desde el paquete mensaje
*/
