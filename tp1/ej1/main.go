package main

import (
	"fmt"
	"net/http"
)

func handleAbout(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w,
		`<!DOCTYPE html>
		<html>
			<head><title>Servidor de Tomi</title></head>
			<body>
				<h1>¡Hola!</h1>
				<p>Corriendo en puerto 8080.</p>
				<a href="/"> Ir a principal </a>
			</body>
		</html>`)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	//1. Manejo de rutas no existentes
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w,
		`<!DOCTYPE html>
		<html>
			<head><title>Hola mundo</title></head>
			<body>
				<h1>¡Hola mundo!</h1>
				<p>Servidor funcionando</p>
				<a href="/about"> Ir a informacion del servidor </a>
			</body>
		</html>`)
}

func main() {

	// Ruta para mostrar el formulario
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/about", handleAbout)
	port := ":8080"
	fmt.Printf("Servidor con formulario escuchando en http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Inicia el servidor
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
