package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Necesario para importar los drivers
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:RGSystem320!.@tcp(172.23.0.1:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexión")
		panic(err) // El panic corta el programa por lo que no requiere un return
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}() // Función anonima que se va a ejecutar automaticamente.

//username:password@tcp(ruta(localhost):puerto(3306))/database --> Estructura de ruta de conexión.
const url = "root:RGSystem320!.@tcp(172.23.0.1:3306)/goweb_db"

// Variable que guarda la conexión:
var db *sql.DB

// Función para conectarme con mi base de datos:
func Connect() {

	conection, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	db = conection

	fmt.Println("Conexión exitosa!!!")
}

// Función para cerrar la conexión con mi base de datos:
func Close() {
	db.Close()
}

// Verificar la conexión:
func Ping() {
	/*
		err := db.Ping() // Si no hay conexión establecida va a devolver un error. Si hay conexión no va a devolver nada.

		if err != nil {
			panic(err)
		}
		Lo mismo se puede hacer así:
	*/

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// Función para crear un tabla en mi motor (MySQL):
func CreateTable(schema string, name string) {

	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("La tabla que desea crear ya existe. Por lo que no podemos volver a crearla!")
	}

	/*
		result, err := db.Exec(schema)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(result.LastInsertId())
		fmt.Println("Esquema creado con Exito!!!")
	*/
}

// Función para verificar si una tabla existe o no:
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)

	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next() // Si puede recorrer la table retorna true y si no puede recorrer la tabla retorna false
}

// Función para hacer Truncate Table: Eliminar todos los registros de una tabla o reiniciar los registros de una tabla que es lo mismo:

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s;", tableName)
	Exec(sql)
}

// Polimorfismo de Exec (db.Exec() de la librería "database/sql"):
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo de Query (db.Query() de la librería "database/sql"):
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows, err

}
