package main

import (
	"database/sql"
	"fmt"
	//"net/http"
	_ "github.com/lib/pq"
)

type userRepository struct {
	db *sql.DB
}

//estructura para mapear de tablas tipo usuario de sql a estructuras de go.
type User struct {
	ID int64
	name string
	email string
}

func abirDB() (*sql.DB, error) {
	//<usuario>:<contraseña>@<nombre_servicio>:<puerto>/<nombre_base_de_datos>?sslmode=disable
	db, err := sql.Open("postgres", "postgres://tomi:123@db:5432/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	return db, nil
}

func (r *userRepository) getUserByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	//creamos conteenedor de tipo User para almacenar los datos traidos de la base de datos.
	user := &User{}
	//asigna lo traido a row hacia el contenedor User.
	err := row.Scan(&user.ID, &user.name, &user.email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) createUser(user *User) (int64, error) {
	row := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.name, user.email)
	//user := &User{}
	var id int64
	//asignamos lo devuelto por la consulta a la variable id.
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func main() {
	db, err := abirDB()
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return
	}
	defer db.Close()

	repo := &userRepository{db: db}

	id, err := repo.createUser(&User{name: "tomi", email: "tomi@example.com"})
	if err != nil {
		fmt.Println("Error al crear el usuario:", err)
		return
	}
	fmt.Println("Usuario creado con ID:", id)

	user, err := repo.getUserByID(1)
	if err != nil {
		fmt.Println("Error al obtener el usuario:", err)
		return
	}
	fmt.Println("Usuario obtenido:", user.name, user.email)
}