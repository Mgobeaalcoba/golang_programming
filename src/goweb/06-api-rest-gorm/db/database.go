package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Necesario para importar los drivers
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Dirección de nuestra base de datos MySQL:

var dsn = "root:RGSystem320!.@tcp(172.23.0.1:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

// Conexión con nuestra base de datos MySQL:

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexión")
		panic(err) // El panic corta el programa por lo que no requiere un return
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}() // Función anonima que se va a ejecutar automaticamente.
