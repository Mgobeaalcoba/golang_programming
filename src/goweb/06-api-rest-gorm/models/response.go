package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string      // De tipo privado por lo que no lleva alias
	respWrite   http.ResponseWriter
}

// Función para crear una respuesta al cliente por default:

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

// Metodo para manejar todo lo que va a responder al cliente.

func (resp *Response) send() {
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(output))
}

// Función para devolver al cliente que vamos a utilizar desde el/los handler/s:

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw) // Respuesta que vamos a devolver al cliente por default
	response.Data = data
	response.send()
}

// Metodo por si ocurre algún error:

func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource Not Found"
}

// Función para devolver al cliente el not found

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.send()
}

// Metodo para entidad no procesable:

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity Not Found"
}

// Función para devolver al cliente la entidad no procesable:

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.send()
}
