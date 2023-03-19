package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Función para responder al cliente si obtenemos data de nuestra base de datos:

func sendData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "aplication/json") // Cambio el content type:
	rw.WriteHeader(status)                             // Escribo el status que recibo en mi función.

	output, _ := json.Marshal(&data) // Transformo en json pero viene en byte. Por lo que debo pasarlo a string abajo.
	fmt.Fprintln(rw, string(output))
}

// Función para responder al cliente en caso de error:

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Fprintln(rw, "Resource Not Found")
}
