package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras o "Clases":

type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

// Variable global Templates para utilizar luego con "Execute templates":

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

// Variable global para acceder a mis templates de Error:

var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Creo mi "Handler Error":

func handlerError(rw http.ResponseWriter, status int) {

	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)

}

// Función que maneje el renderizamiento de nuestros templates:

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		handlerError(rw, http.StatusInternalServerError)
	}

}

// Funciones:

func Saludar() string {
	return "Hola desde una función"
}

// Creo mi Handler para mi Mux principal:

func Index(rw http.ResponseWriter, r *http.Request) {

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}

	courseList := []Curso{c1, c2, c3, c4}

	usuario1 := Usuarios{"Mariano", 35, true, true, courseList}

	renderTemplate(rw, "index.html", usuario1)

}

// Segundo Handler construido para este ejercicio:

func Registro(rw http.ResponseWriter, r *http.Request) {

	renderTemplate(rw, "registro.html", nil)

}

func main() {

	// Archivos estátivos:
	staticFile := http.FileServer(http.Dir("static"))

	// Creamos nuestro Mux:

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Mux de staticFiles:

	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	// Creamos nuestro Servidor:

	server := &http.Server{
		Addr:    "localhost:5006",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 5006")
	fmt.Println("Run server: http://localhost:5006")

	log.Fatal(server.ListenAndServe())

}
