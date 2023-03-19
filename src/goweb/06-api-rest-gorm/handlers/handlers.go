package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Importante pasar a GORM los objetos con puntero (Lo objeto mismo y no una copia) para que funcione de forma correcta!!!

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	// Devuelve todos los registros: por lo que si no tenemos ninguno no es necesario manejar errores sino que devolvemos la lista vacia:

	users := models.Users{} // Creo una estructura vacia de users.

	// Con GORM obtener todos los datos de nuestra base de datos es mucho mas facil aún ya que no necesitamos querys...

	db.Database.Find(&users) // En users debe guardar todos los datos de la tabla users

	sendData(rw, users, http.StatusOK) // Armamos la respuesta para el cliente.

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}

}

func getUserById(r *http.Request) (models.User, *gorm.DB) {

	//Obtener ID del request del cliente:
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}
	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user) // Metodo de GORM que guarda un nuevo registro en nuestra base de datos.
		sendData(rw, user, http.StatusOK)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro:
	var userId int64

	if user_ant, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound) // No encontrar un registro es un "Status Not Found"
	} else {
		userId = user_ant.Id

		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user) // Metodo de GORM que guarda un nuevo registro en nuestra base de datos.
			sendData(rw, user, http.StatusOK)
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro:

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound) // No encontrar un registro es un "Status Not Found"
	} else {
		sendData(rw, user, http.StatusOK)
		db.Database.Delete(&user)
	}
}
