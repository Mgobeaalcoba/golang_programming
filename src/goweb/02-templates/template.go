package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras o "Clases"
type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso // Arreglo de tipo Curso que voy a crear debajo
}

type Curso struct {
	Nombre string
}

// Funciones:

func Saludar() string {
	return "Hola desde una función"
} // Esta función puede ser ejecutada desde nuestro archivo renderizado HTML.

// Creo mi Handler para mi Mux principal:
func Index(rw http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}

	courseList := []Curso{c1, c2, c3, c4} // Este arreglo que lo mandaremos como atributo del usuario1 va a ser necesario iterarlo desde nuestro template HTML.

	// ¿Que vamos a retornarle a la petición del cliente?
	// fmt.Fprintln(rw, "Hola Mundo!!!") // Como respuesta al cliente luego de su petición enviamos por pantalla un "Hola Mundo"

	template, err := template.ParseFiles("index.html") // Ruta del archivo HTML con el que me quiero conectar para luego renderizarlo.

	usuario1 := Usuarios{"Mariano", 35, true, true, courseList}

	if err != nil {
		panic(err)
	} else {
		//template.Execute(rw, nil) // El segundo parametro es una data que ahora no vamos a enviar.

		template.Execute(rw, usuario1) // Mando mi objeto usuario como data a mi template HTML para usarlo y mostrarlo alla.
	}

	// El "Hola Mundo" que se imprime ahora en mi dominio viene desde el HTML en lugar de venir desde mi template.go

}

func main() {
	// Creamos nuestro Mux:

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Creamos nuestro Servidor:

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 5000")
	fmt.Println("Run server: http://localhost:5000")

	log.Fatal(server.ListenAndServe())
}

/*
Para renderizar un archivo html vamos a usar la librería "html/template"

Renderizar en términos simples significa tomar un archivo HTML (o una plantilla HTML) y rellenar los espacios en blanco con contenido dinámico para producir una página web completa.

En el contexto de Golang, la librería "Template" es una herramienta útil para renderizar archivos HTML. Esta librería te permite definir plantillas HTML con "placeholders" (lugares para el contenido dinámico) y luego llenar esos placeholders con valores dinámicos en tiempo de ejecución.

Por ejemplo, si tienes una plantilla HTML que muestra información de un usuario, podrías utilizar la librería "Template" para rellenar esa plantilla con la información del usuario específico que estás mostrando.

En resumen, la librería "Template" en Golang te ayuda a renderizar archivos HTML de manera dinámica, llenando los espacios en blanco con contenido específico para cada situación.

*/
