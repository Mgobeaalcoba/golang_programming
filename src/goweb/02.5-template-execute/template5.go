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

// Variable global Templates para utilizar luego con "Execute templates"

var templates = template.Must(template.New("T").ParseFiles("index.html", "base.html")) // A esta variable hay que definirla con var antes e "=" luego. No de forma resumida (":=") // En ParseFiles se cargan todos los templates que luego voy a poder usar con ExecuteTemplate // En el .New puedo poner cualquier cosa. La T es un ejemplo.

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

	// template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html")) // Comento mi template individual para armar una variable Templates general afuera del Handler

	usuario1 := Usuarios{"Mariano", 35, true, true, courseList}

	// template.Execute(rw, usuario1) // Ya no lo vamos a usar así!!!

	err := templates.ExecuteTemplate(rw, "index.html", usuario1)

	if err != nil {
		panic(err)
	}
}

func main() {
	// Creamos nuestro Mux:

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Creamos nuestro Servidor:

	server := &http.Server{
		Addr:    "localhost:5004",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 5004")
	fmt.Println("Run server: http://localhost:5004")

	log.Fatal(server.ListenAndServe())
}

/*
Para renderizar un archivo html vamos a usar la librería "html/template"

Renderizar en términos simples significa tomar un archivo HTML (o una plantilla HTML) y rellenar los espacios en blanco con contenido dinámico para producir una página web completa.

En el contexto de Golang, la librería "Template" es una herramienta útil para renderizar archivos HTML. Esta librería te permite definir plantillas HTML con "placeholders" (lugares para el contenido dinámico) y luego llenar esos placeholders con valores dinámicos en tiempo de ejecución.

Por ejemplo, si tienes una plantilla HTML que muestra información de un usuario, podrías utilizar la librería "Template" para rellenar esa plantilla con la información del usuario específico que estás mostrando.

En resumen, la librería "Template" en Golang te ayuda a renderizar archivos HTML de manera dinámica, llenando los espacios en blanco con contenido específico para cada situación.

*/
