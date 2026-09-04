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

func (r *userRepository) listUsers() ([]*User, error) {
    //se usa Query en lugar de QueryRow para obtener multiples filas.
    rows, err := r.db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    //cerramos manualmente la conexion
    defer rows.Close()

    users := []*User{} //creamos slides de usuarios vacio
    
	//recorremos todas las filas
    for rows.Next() {
        i := &User{} //contenedor individual de tipo User
        if err := rows.Scan(&i.ID, &i.name, &i.email); err != nil {
            return nil, err
        }
        users = append(users, i)
    }

    //verificamos si hubo algun error al recorrer las filas
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *userRepository) updateUser(id int64, name string, email string) (string, error) {
	row := r.db.QueryRow("update users set name = $2, email = $3 where id = $1 returning name", id, name, email)
	var nombre string
	err := row.Scan(&nombre)
	if err != nil {
		return "", err
	}

	return nombre, nil
}

func (r *userRepository) deleteUser(id int64) error {
	//se usa Exec ya que no se espera un resultado.
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}


func main() {
	db, err := abirDB()
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
	}
	defer db.Close()

	repo := &userRepository{db: db}

	//creacion usuario
	id, err := repo.createUser(&User{name: "tomi", email: "tomi@example.com"})
	if err != nil {
		fmt.Println("Error al crear el usuario:", err)
	}
	fmt.Println("Usuario creado con ID:", id)

	//obtener usuario por su ID
	user, err := repo.getUserByID(1)
	if err != nil {
		fmt.Println("Error al obtener el usuario:", err)
	}
	fmt.Println("Usuario obtenido:", user.name, user.email)

	//listar usuarios
	users, err := repo.listUsers()
	if err != nil {
		fmt.Println("Error al listar los usuarios:", err)
	}
	for _, user := range users {
		fmt.Println("Usuario:", user.name, user.email)
	}

	//actualizar usuario
	updatedName, err := repo.updateUser(1, "nuevoNombre", "nuevoEmail@example.com")
	if err != nil {
		fmt.Println("Error al actualizar el usuario:", err)
	}
	fmt.Println("Usuario actualizado con nuevo nombre:", updatedName)

	//eliminar usuario
	err = repo.deleteUser(1)
	if err != nil {
		fmt.Println("Error al eliminar el usuario:", err)
	}
	fmt.Println("Usuario eliminado")

}