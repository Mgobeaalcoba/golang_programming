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

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html")) // Voy a reemplazar ParseFiles por ParseGlob para indicar con un solo direccionamiento todo mi directorio de templates.

// Variable global para acceder a mis templates de Error:
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Creo mi "Handler Error" Handler para mostrar con un template especifico cuando ocurra un error en nuestra aplicación.

func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status) // Modifica el status del template que estamos mostrando con el que corresponde al error a mostrar.
	errorTemplate.Execute(rw, nil)
} // Creada mi función para llamar a un template cuando deba manejar errores entonces ahora voy a editar mi función para manejo de errores "renderTemplate"

// Función que maneje el renderizamiento de nuestros templates, lo cual, entre otras cosas implica que maneje los posibles errores. Así no los tengo que estar manejando de forma individual en cada Handler:

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		// Si tenemos un error y solo usamos el panic(err) lo unico que logramos es mostrar al usuario el motivo de error pero rompemos la navegación en la pagina. Si queremos evitar esto entonces debemos manejar los erorres:
		// http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		// Ahora no romperemos la navegación pero tmb le mostraremos al usuario que estamos frente a un error.
		// Uso handlerError como acción frente a un error, el cual me va a cargar mi template "error.html":
		handlerError(rw, http.StatusInternalServerError)
		// Ahora con esto tenemos nuestro propio template para manejar errores.
		// Pero tenemos un problema aún que es al cargar de forma exitosa un nuevo template el status que nos mostrará será "200" cuando debería ser status "500". Por eso debemos enviar el status que queremos mostrar a la función handlerError
	}
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

	// template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html")) // Comento mi template individual para armar una variable Templates general afuera del Handler

	usuario1 := Usuarios{"Mariano", 35, true, true, courseList}

	renderTemplate(rw, "index.html", usuario1) // Pongo inde.html a proposito para romper y probar el manejo de errores. Debería ir "index.html"

	// template.Execute(rw, usuario1) // Ya no lo vamos a usar así!!!

	/*

		Este trabajo que se está repitiendo en cada Handler se lo voy a pasar mi función "RenderTemplate":

		err := templates.ExecuteTemplate(rw, "index.html", usuario1)

		if err != nil {
			panic(err)
		}
	*/
}

// Segundo Handler construido para este ejercicio:

func Registro(rw http.ResponseWriter, r *http.Request) {

	// Voy a usar mi función "renderTemplate para validar si hay un error o no. Ya no lo voy a hacer en cada Handler"

	renderTemplate(rw, "registro.html", nil)

	/*
		err := templates.ExecuteTemplate(rw, "registro.html", nil)

		if err != nil {
			panic(err)
		}
	*/
}

// Cada Handler debe corresponderse con un nuevo archivo HTML al cual responde. En este caso debo crear un archivo en templates llamaddo "registro.html"

// Tmb debo agregar un nuevo "mux" para esta nueva sección que estoy creando:

func main() {
	// Creamos nuestro Mux:

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Creamos nuestro Servidor:

	server := &http.Server{
		Addr:    "localhost:5005",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 5005")
	fmt.Println("Run server: http://localhost:5005")

	log.Fatal(server.ListenAndServe())
}

/*
Para renderizar un archivo html vamos a usar la librería "html/template"

Renderizar en términos simples significa tomar un archivo HTML (o una plantilla HTML) y rellenar los espacios en blanco con contenido dinámico para producir una página web completa.

En el contexto de Golang, la librería "Template" es una herramienta útil para renderizar archivos HTML. Esta librería te permite definir plantillas HTML con "placeholders" (lugares para el contenido dinámico) y luego llenar esos placeholders con valores dinámicos en tiempo de ejecución.

Por ejemplo, si tienes una plantilla HTML que muestra información de un usuario, podrías utilizar la librería "Template" para rellenar esa plantilla con la información del usuario específico que estás mostrando.

En resumen, la librería "Template" en Golang te ayuda a renderizar archivos HTML de manera dinámica, llenando los espacios en blanco con contenido específico para cada situación.

*/
