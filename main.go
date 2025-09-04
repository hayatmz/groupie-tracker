package main

import (
	"fmt"
	web "groupietrack/web"
	"net/http"
)

const port = ":4444"

// Sert le serveur sur le port de la constante.
func main() {
	http.HandleFunc("/", web.ErrorHandler)
	http.HandleFunc("/artist", web.ArtistHandler)
	http.HandleFunc("/results-search", web.SearchHandler)
	fmt.Println("(http://localhost:4444) - Server started on port", port)
	http.ListenAndServe(port, nil)
}