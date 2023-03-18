package handlers

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(rw)
		//models.SendData(rw, users)
	} else {
		//models.SendNotFound(rw) // La data es users
		models.SendData(rw, users)
	}

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	user := models.User{}              // Struct vacio que voy a completar.
	decoder := json.NewDecoder(r.Body) // Obtiene el cuerpo del request y lo decofica de tipo json a tipo objeto

	// El metodo Decode carga los datos del request body al objeto construido vacio. Si lo logra no devuelve nada, pero si no lo logra devuelve un error que debemos manejar.
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save() // Metodo de user que verifica si el id existe o no y en función de ello registra uno nuevo o edita.
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		decoder := json.NewDecoder(r.Body) // Obtiene el cuerpo del request y lo decofica de tipo json a tipo objeto

		// El metodo Decode carga los datos del request body al objeto user traido de mi base de datos. Si lo logra no devuelve nada, pero si no lo logra devuelve un error que debemos manejar.
		if err := decoder.Decode(&user); err != nil {
			models.SendUnprocessableEntity(rw)
		} else {
			user.Save() // Metodo de user que verifica si el id existe o no y en función de ello registra uno nuevo o edita.
			models.SendData(rw, user)
		}
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		models.SendData(rw, user)
		user.Delete() // Uso el metodo Delete para eliminarlo de mi base de datos.
	}
}

func getUserByRequest(r *http.Request) (*models.User, error) {
	//Obtener ID del request del cliente:

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	if user, err := models.GetUser(userId); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
