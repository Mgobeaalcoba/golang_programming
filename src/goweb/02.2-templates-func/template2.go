package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Funciones:

func Saludar(nombre string) string {
	return "Hola desde una función " + nombre
} // Esta función puede ser ejecutada desde nuestro archivo renderizado HTML.

// Creo mi Handler para mi Mux principal:
func Index(rw http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	} // Creo mi mapa de funciones que voy a enviar a mi HTML.

	template, err := template.New("index.html").Funcs(funciones).
		ParseFiles("index.html") // En Funcs se envia una mapa de funciones que debo crear antes.

	if err != nil {
		panic(err)
	} else {
		//template.Execute(rw, nil) // El segundo parametro es una data que ahora no vamos a enviar.

		template.Execute(rw, nil) // Mando mi objeto usuario como data a mi template HTML para usarlo y mostrarlo alla.
	}

	// El "Hola Mundo" que se imprime ahora en mi dominio viene desde el HTML en lugar de venir desde mi template.go

}

func main() {
	// Creamos nuestro Mux:

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Creamos nuestro Servidor:

	server := &http.Server{
		Addr:    "localhost:5001",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 5001")
	fmt.Println("Run server: http://localhost:5001")

	log.Fatal(server.ListenAndServe())
}

/*
Para renderizar un archivo html vamos a usar la librería "html/template"

Renderizar en términos simples significa tomar un archivo HTML (o una plantilla HTML) y rellenar los espacios en blanco con contenido dinámico para producir una página web completa.

En el contexto de Golang, la librería "Template" es una herramienta útil para renderizar archivos HTML. Esta librería te permite definir plantillas HTML con "placeholders" (lugares para el contenido dinámico) y luego llenar esos placeholders con valores dinámicos en tiempo de ejecución.

Por ejemplo, si tienes una plantilla HTML que muestra información de un usuario, podrías utilizar la librería "Template" para rellenar esa plantilla con la información del usuario específico que estás mostrando.

En resumen, la librería "Template" en Golang te ayuda a renderizar archivos HTML de manera dinámica, llenando los espacios en blanco con contenido específico para cada situación.

*/
