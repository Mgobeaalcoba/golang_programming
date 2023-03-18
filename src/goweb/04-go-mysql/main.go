package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//fmt.Println("Verificamos que estemos conectados con la database: ")
	db.Ping()
	fmt.Println("Verificación de conexión Ping exitosa")
	//db.CreateTable(models.UserSchema, "users") //Guardado el schema en users.go y creada la función CreateTable en database.go ya puedo ejecutar la misma pasandole mi esquema desde main.go
	// db.Ping()                         //No debería devolver nada
	//fmt.Println("Verifico si existe users:")
	//fmt.Println(db.ExistsTable("users"))
	//fmt.Println("Verifico si existe lalala:")
	//fmt.Println(db.ExistsTable("lalala"))

	db.TruncateTable("users") // Si no hay error funciona correctamente.

	//Creación de usuario
	user1 := models.CreateUser("MarianoGobea", "Pikachu123", "gobeamariano@gmail.com")
	fmt.Println(user1)

	user2 := models.CreateUser("NicoleReina", "TeAmo123", "nicolediosa@gmail.com")
	fmt.Println(user2)

	user3 := models.CreateUser("Lisandrito", "Capo123", "lichitacapo@gmail.com")
	fmt.Println(user3)

	user4 := models.CreateUser("Lauti", "Autiiiii123", "chulolauti@gmail.com")
	fmt.Println(user4)

	users := models.ListUsers()
	fmt.Println(users)

	fmt.Println("Voy a buscar al usuario id=4")
	userBuscado := models.GetUser(4) // Traigo el id 4
	fmt.Println(userBuscado)
	userBuscado.Id = 0
	userBuscado.Save() // Debería guardar un registro numero 5 con los datos originales de Lauti

	fmt.Println("Veamos nuevamente la lista completa a ver si tenemos 5 registros o 4 registros:")
	users = models.ListUsers()
	fmt.Println(users)

	userBuscado.Id = 4
	userBuscado.UserName = "LautiGobea"
	userBuscado.Password = "HuggyWuggy"
	userBuscado.Email = "LautiTeAmo@gmail.com"
	userBuscado.Save() // Envio a save el user 4 y como no es numero 0 va a actualizar sus datos.

	fmt.Println("Imprimo el actualizado solamente:")
	fmt.Println(userBuscado) //Imprimo el actualizado.

	fmt.Println("Imprimo la lista completa de nuevo:")
	users = models.ListUsers()
	fmt.Println(users)

	fmt.Println("Borro el registro 4 y vuelvo a imprimir la lista:")
	userBuscado.Delete()
	users = models.ListUsers()
	fmt.Println(users)

	db.Close()
	//db.Ping() //Debería deovolver el error xq la conexión ya se cerró.

}
