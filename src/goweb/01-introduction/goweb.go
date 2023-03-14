package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers o controlador que maneja una solicitud http particular

func Hola(rw http.ResponseWriter, r *http.Request) {
	// Metodos: Get, Post, Put, etc.
	// Para saber con cual metodo estamos trabajando debemos hacerlo con una propiedad del Request llamada Method
	fmt.Println("El metodo es + ", r.Method) // Sale por consola
	fmt.Fprintln(rw, "Hola Mundo")           // Sale impreso en la web
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	//http.Error(rw, "La Pagina no funciona", 404)
	http.Error(rw, "La Pagina no funciona", http.StatusNotFound) // Go recomienda usar las constantes de http en lugar de los valores de los status de las respuestas directamente como en el codigo comentado.
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)          // Devuelve la URL completa
	fmt.Println(r.URL.RawQuery) // Devuelve la parte de la URL donde están los datos que paso. Luego del "?"
	fmt.Println(r.URL.Query())  // Convierte la información que muestra el RawQuery en una mapa para poder mostrarla y trabajarla de forma mas prolija.

	// Ahora que ya obtengo en mapa los parametros que el usuario me pasa puedo guardarlos en variables que me permitan luego reutilizarlos. ¿Como?

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Printf("El nombre es %v y su edad es %v\n", name, age)              // Imprime por consola los datos recibidos en la URL
	fmt.Fprintf(rw, "Hola mi buen amigo %s, tu edad es %s !!\n", name, age) // Imprimo por pantalla del navegador un saludo para el usuario que ingresó sus datos. Uso "%s" porque todo lo que recibo por Query son strings por default.

	http.Error(rw, "Este es un error", http.StatusConflict)
}

func main() {

	// Mux o Multiplexor:
	mux := http.NewServeMux()
	//mux.HandleFunc() // Usamos el HandleFunc desde el Mux en lugar de desde http. Parece lo mismo pero es totalmente distinto.
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// Router

	// http.HandleFunc("/", Hola) // Función que se ejecuta al acceder a determinada sección de nuestro servidor. En este caso al home.

	// Para probar los status de las respuestas http vamos a crear un nuevo router y un nuevo handler para manejarlo
	// http.HandleFunc("/page", PaginaNF)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	//Mensajes que van a aparecer por consola cuando encienda el servidor.

	// Crear mi propio Servidor:

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 8000")
	fmt.Println("Run server: http://localhost:8000")

	// Crear servidor:

	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//log.Fatal(http.ListenAndServe("localhost:8000", mux))

	log.Fatal(server.ListenAndServe())

	// Los parametro son dirección donde estará hosteado nuestro servidor y un "mux" que por el momento no enviamos y pasamos un "nil". Nil implica que nuestro servidor está trabajando con mux automaticos o mux por defecto.

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

	/*
		Muchos lenguajes de programación ofrecen la posibilidad de reiniciar los servidores de forma automatica cada vez que se guardan cambios en el mismo. Golang también lo ofrece a traves de su librería fresh

		Link: https://github.com/gravityblast/fresh

		Como prerequisito para que funcione fresh debo tener mi archivo go.mod usando el comando:

		go mod init "nombre" (en este caso goweb)

		Luego puedo instalar fresh y ejecutarlo

		En linux no se puede ejecutar con el comando fresh directamente

		Debo realizar lo siguiente:

		1- go run github.com/pilu/fresh@latest
	*/
}
