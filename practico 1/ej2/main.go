package main

import (
	"fmt"
	"net/http"
)

// HTML del formulario
const loginForm = `
	<!DOCTYPE html>
	<html>
		<head><title>Contacto</title></head>
		<body>
			<h2>Login</h2>
			<form action="/contacto" method="POST">
				<label>Nombre:</label>
				<input type="text" name="nombre"><br>
				<label>mail:</label>
				<input type="email" name="mail"><br>
				<label>Mensaje:</label>
				<input type="text" name="mensaje"><br>
				<button type="submit">Enviar</button>
			</form>
		</body>
	</html>`

func handleContact(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/contacto" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, loginForm)
		return
	}

	if r.Method == http.MethodPost {

		// 1. Parsear los datos del formulario (¡Crucial!)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error al parsear", http.StatusBadRequest)
			return
		}
		username := r.FormValue("nombre")
		mail := r.FormValue("mail")
		mnsj := r.FormValue("mensaje")

		// 2. Validamos que no estén vacíos
		if username == "" || mail == "" || mnsj == "" {
			http.Error(w, "Error: Campo vacio", http.StatusBadRequest)
			return
		}

		// 3. Generar y enviar la respuesta HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w,
			`<!DOCTYPE html>
			<html>
				<head><title>Bienvenido</title></head>
				<body>
					<h1>¡Hola, %s!</h1>
					<p>Recibimos tus datos.</p>
					<a href="/contacto">Volver</a>
				</body>
			</html>`, username)
		return
	}
}

func main() {
	// Ruta para mostrar el formulario
	http.HandleFunc("/contacto", handleContact)
	// Ruta para procesar el formulario
	//http.HandleFunc("/login", handleLogin)
	port := ":8080"
	fmt.Printf("Servidor con formulario escuchando en http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Inicia el servidor
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
