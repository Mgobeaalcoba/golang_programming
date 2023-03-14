package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hola Mundo")
}

func main() {

	// Router
	http.HandleFunc("/", Hola) // Función que se ejecuta al acceder a determinada sección de nuestro servidor. En este caso al home.

	//Mensajes que van a aparecer por consola cuando encienda el servidor.

	fmt.Println("El servidor está corriendo en el puerto 8000")
	fmt.Println("Run server: http://localhost:8000")

	// Crear servidor:

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

	// Los parametro son dirección donde estará hosteado nuestro servidor y un mocks que por el momento no enviamos y pasamos un "nil"

	// Solo con el codigo de arriba ya tendríamos nuestro servidor creado.

	// ¿Como lo ejecutamos?
	// go run "nombre.go"
	// Quedará la consola trabajando ya que el servidor está encendido.

	/*
		Luego de ello ya puedo cargar mi pagina en la dirección donde la hostee a traves de un navegador poniendo en el ejemplo de arriba:

		localhost:8000

		Nos devolverá una pagina vacia con la leyenda:

		404 page not found

		mientras la misma no tenga contenido.
	*/
}
