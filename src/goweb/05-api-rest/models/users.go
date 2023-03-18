package models

import "apirest/db"

type User struct {
	Id       int64  `json:"id"` //Alias para json
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `
CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50) NOT NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
`

// Construir usuario: Voy a armar un constructor de la clase/struct user:

func NewUser(username, password, email string) *User {
	user := &User{
		UserName: username,
		Password: password,
		Email:    email,
	}
	return user
}

// Crear usuario e insertar a la base de datos:
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Insertar Registro (Es un metodo de la struct User)
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?;"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	// Guardo el result para poder obtener el ultimo id y guardarlo en mi objeto con el mismo id que en mi base de datos:
	user.Id, _ = result.LastInsertId()

}

// Actualizar registro (Es un metodo de la struct User):
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?;"
	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)
}

// Insertar o actualizar (guardar o editar) registro (Es un metodo de la struct User):
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Eliminar UN registro (Es un metodo de la struct User):
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?;"
	db.Exec(sql, user.Id)
}

// Listar TODOS los registros de mi table/base de datos:
func ListUsers() (Users, error) {
	sql := "SELECT id, username, password, email FROM goweb_db.users;"
	users := Users{}
	if rows, err := db.Query(sql); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
			users = append(users, user)
		}
		return users, nil
	}
	/*
		rows, err := db.Query(sql) // Rows es un iterable que se logra iterar con el metodo .Next() de "database/sql"

		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
			users = append(users, user)
		}

		return users, err
	*/
}

// Obtener SOLAMENTE un registro:
func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM goweb_db.users WHERE id=?;" // Filtro con el WHERE de sql para traerme solo el id que me interesa.

	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email) // Le cargo todos los datos al objeto vacio que cree arriba
		}

		return user, nil
	}

}
