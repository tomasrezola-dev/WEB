package main

import (
	"fmt"
	"html"
	"net/http"
)

// HTML del formulario
const loginForm = `
<!DOCTYPE html>
<html>
<head>
	<title>Contacto</title>
</head>
<body>
	<h2>Contacto</h2>

	<form action="/contacto" method="POST">
		<label>Nombre:</label>
		<input type="text" name="nombre"><br><br>

		<label>Mail:</label>
		<input type="email" name="mail"><br><br>

		<label>Mensaje:</label>
		<input type="text" name="mensaje"><br><br>

		<button type="submit">Enviar</button>
	</form>
</body>
</html>
`

func handleContact(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/contacto" {
		http.NotFound(w, r)
		return
	}

	// Parseamos los parámetros del formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al parsear los datos", http.StatusBadRequest)
		return
	}

	//obtencion ip con puerto
	ip := r.RemoteAddr

	// Si es POST, validar los datos
	if r.Method == http.MethodPost {

		nombre := r.FormValue("nombre")
		mail := r.FormValue("mail")
		mensaje := r.FormValue("mensaje")

		if nombre == "" || mail == "" || mensaje == "" {
			http.Error(w, "Error: campo vacío", http.StatusBadRequest)
			return
		}
	}

	// Respuesta HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w,
		`<!DOCTYPE html>
		<html>
			<head>Información de la petición</head>
			<body>
				<h1>Información de la petición HTTP</h1>
				<h2>Método HTTP</h2>
				<p>`)
	fmt.Fprint(w, html.EscapeString(r.Method))
	fmt.Fprint(w,
		`</p>
				<h2>IP del cliente</h2>
				<p>`)
	fmt.Fprint(w, html.EscapeString(ip))
	fmt.Fprint(w,
		`</p>
				<h2>Cabeceras recibidas</h2>
				<ul>`)
	for nombre, valores := range r.Header {
		for _, valor := range valores {
			fmt.Fprintf(w, "<li><strong>%s:</strong> %s</li>",
				html.EscapeString(nombre),
				html.EscapeString(valor))
		}
	}
	fmt.Fprint(w,
		`</ul>
				<h2>Parámetros GET</h2>
				<ul>`)
	for nombre, valores := range r.URL.Query() {
		for _, valor := range valores {
			fmt.Fprintf(w, "<li><strong>%s:</strong> %s</li>",
				html.EscapeString(nombre),
				html.EscapeString(valor))
		}
	}
	fmt.Fprint(w,
		`</ul>
				<h2>Parámetros POST</h2>
				<ul>`)
	for nombre, valores := range r.PostForm {
		for _, valor := range valores {
			fmt.Fprintf(w, "<li><strong>%s:</strong> %s</li>",
				html.EscapeString(nombre),
				html.EscapeString(valor))
		}
	}
	fmt.Fprint(w,
		`</ul>
				<h2>Formulario</h2>`)
	fmt.Fprint(w, loginForm)
	fmt.Fprint(w, `
			</body>
		</html>
	`)
}

func main() {

	http.HandleFunc("/contacto", handleContact)

	port := ":8080"

	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
