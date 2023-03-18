package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	//rw.Header().Set("Content-Type", "text/xml")

	// Obtener todos los usuarios de la base de datos.
	db.Connect()
	users := models.ListUsers() // Nos va a traer el listado en tipo objeto. Pero nosotros debemos devolverlo al usuario en tipo JSON, XML o algún otro tipo de datos globales. Porque así se trabaja con API-REST.
	db.Close()

	output, _ := json.Marshal(users) // Devuelve los valores transformados en tipo Byte y un error. Por lo que debo capturar esos valores, transformarlos a string y devolverlo.
	fmt.Fprintln(rw, string(output))

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	// Obtener ID: Debo capturar el id que envía el cliente mediante url en su consulta al endpoint. Eso lo hacemos con "gorilla/mux"
	vars := mux.Vars(r)                   // Devuelve mapa de tipo string. Pero yo necesito un entero para mi func GetUser() por lo que debo transformarlo primero.
	userId, _ := strconv.Atoi(vars["id"]) // Convierto a entero con strconv.Atoi()

	// Obtener un usuario de mi base de datos
	db.Connect()

	user := models.GetUser(userId)

	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// Crear un registro nuevo en MySQL mediante el cliente pegandole al endpoint creado para tales fines:

	rw.Header().Set("Content-Type", "application/json")

	// En este caso ya no alcanza con obtener solo el id. Sino que debemos obtener todo el registro (salvo el id que lo traigo de la base) que vamos a recibir para pasarselo a la función models.CreateUser()

	user := models.User{} // Struct vacio que voy a completar.

	decoder := json.NewDecoder(r.Body) // Obtiene el cuerpo del request y lo decofica de tipo json a tipo objeto

	// El metodo Decode carga los datos del request body al objeto construido vacio. Si lo logra no devuelve nada, pero si no lo logra devuelve un error que debemos manejar.
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save() // Metodo de user que verifica si el id existe o no y en función de ello registra uno nuevo o edita.
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output)) // Devolvemos igual el usuario creado. Los parametros de la creación se deben pasar por el "Body" de la consulta al endpoint
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// Actualizar un registro existente y que se pasa por request en MySQL mediante el cliente pegandole al endpoint creado para tales fines:

	rw.Header().Set("Content-Type", "application/json")

	// Obtener ID: Debo capturar el id que envía el cliente mediante url (ya que ese es el usuario que me tengo que traer para editar) en su consulta al endpoint. Eso lo hacemos con "gorilla/mux"
	vars := mux.Vars(r)                   // Devuelve mapa de tipo string. Pero yo necesito un entero para mi func GetUser() por lo que debo transformarlo primero.
	userId, _ := strconv.Atoi(vars["id"]) // Convierto a entero con strconv.Atoi()

	// Obtener un usuario de mi base de datos
	db.Connect()

	user := models.GetUser(userId) // Me traigo al usuario que pidió el cliente

	db.Close()

	decoder := json.NewDecoder(r.Body) // Obtiene el cuerpo del request y lo decofica de tipo json a tipo objeto

	// El metodo Decode carga los datos del request body al objeto user traido de mi base de datos. Si lo logra no devuelve nada, pero si no lo logra devuelve un error que debemos manejar.
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save() // Metodo de user que verifica si el id existe o no y en función de ello registra uno nuevo o edita.
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output)) // Devolvemos igual el usuario creado. Los parametros de la creación se deben pasar por el "Body" de la consulta al endpoint
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener ID: Debo capturar el id que envía el cliente mediante url (ya que ese es el usuario que me tengo que traer para editar) en su consulta al endpoint. Eso lo hacemos con "gorilla/mux"
	vars := mux.Vars(r)                   // Devuelve mapa de tipo string. Pero yo necesito un entero para mi func GetUser() por lo que debo transformarlo primero.
	userId, _ := strconv.Atoi(vars["id"]) // Convierto a entero con strconv.Atoi()

	// Obtener un usuario de mi base de datos
	db.Connect()

	user := models.GetUser(userId) // Me traigo al usuario que pidió el cliente
	user.Delete()                  // Uso el metodo Delete para eliminarlo de mi base de datos.

	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output)) // Devolvemos igual el usuario creado. Los parametros de la creación se deben pasar por el "Body" de la consulta al endpoint
}
