package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 1. envoltorio para la respuesta
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// 2. sobreescribos el metodo writer para que se redireccione hacia
// compresor GZIP
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// 3. creacion de middleware
func middlewareGzip(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//verificamos si el navegador soporta GZIP.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			//si no, realizamos normalmente la respuesta con el file server
			next.ServeHTTP(w, r)
		}
		//si soporta, lo marcamos en el encabezado.
		w.Header().Set("Content-Encoding", "gzip")
		//creacion compresor
		gz := gzip.NewWriter(w)
		//el compresor se cierra al terminar de enviar datos.
		defer gz.Close()
		//intercabiamos el escritor normal por el de GZIP
		gzw := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		//se le pasa la peticion al file server pero este lo envia
		//al compresor que lo envia al navegador.
		next.ServeHTTP(gzw, r)
	})
}

func main() {
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))

	//http.Handle("/", middlewareGzip(fileServer))
	http.Handle("/", fileServer)
	port := ":8080"
	fmt.Printf("Servidor con formulario escuchando en http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Inicia el servidor
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
