package main

import (
	"log"
	"net/http"
	"templ-example/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
