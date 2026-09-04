package main

import (
	"testing"
)

func TestCreateUser(t *testing.T) {

	db, err := abirDB()
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	defer db.Close()

	repo := &userRepository{db: db}

	//1. creacion usuario
	id, err := repo.createUser(&User{name: "testuser", email: "testuser@example.com"})
	if err != nil {
		t.Errorf("Error al crear el usuario: %v", err)
	}

	//2. obtener usuario por su ID
	user, err := repo.getUserByID(id)
	if err != nil {
		t.Errorf("Error al obtener el usuario: %v", err)
	}

	//3. actualizar usuario y comprobar cambio
	updatedName, err := repo.updateUser(1, "nuevoNombre", "nuevoEmail@example.com")
	if err != nil {
		t.Errorf("Error al actualizar el usuario: %v", err)
	}
	user, err = repo.getUserByID(1)
	if err != nil {
		t.Errorf("Error al obtener el usuario: %v", err)
	}
	if user.name != updatedName {
		t.Errorf("El nombre del usuario no se actualizó correctamente. Esperado: %s, Obtenido: %s", updatedName, user.name)
	}

	//4. listar usuarios y comprobar que el usuario creado está en la lista
	users, err := repo.listUsers()
	if err != nil {
		t.Errorf("Error al listar los usuarios: %v", err)
	}
	found := false
	for _, u := range users {
		if u.ID == user.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("El usuario creado no se encuentra en la lista de usuarios")
	}

	//5. ELiminar usuario y comprobar que ya no está en la lista
	err = repo.deleteUser(1)
	if err != nil {
		t.Errorf("Error al eliminar el usuario: %v", err)
	}
	user, err = repo.getUserByID(1)
	if err == nil {
		t.Errorf("Se obtuvo el usuario pero se esperaba que no existiera después de la eliminación")
	}
	
}
