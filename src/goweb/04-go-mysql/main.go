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
	db.CreateTable(models.UserSchema, "users") //Guardado el schema en users.go y creada la función CreateTable en database.go ya puedo ejecutar la misma pasandole mi esquema desde main.go
	// db.Ping()                         //No debería devolver nada
	fmt.Println("Verifico si existe users:")
	fmt.Println(db.ExistsTable("users"))
	fmt.Println("Verifico si existe lalala:")
	fmt.Println(db.ExistsTable("lalala"))
	db.Close()
	//db.Ping() //Debería deovolver el error xq la conexión ya se cerró.

}
