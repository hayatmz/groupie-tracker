package main

import (
	"fmt"
	web "groupietrack/web"
	"net/http"
	"os"
)

const port = ":4444"

// Sert le serveur sur le port de la constante.
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4444" // fallback local
	}

	// Servir les fichiers statiques (CSS, img)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", web.ErrorHandler)
	http.HandleFunc("/artist", web.ArtistHandler)
	http.HandleFunc("/results-search", web.SearchHandler)

	fmt.Println("(http://localhost:" + port + ") - Server started")
	http.ListenAndServe(":"+port, nil)
}