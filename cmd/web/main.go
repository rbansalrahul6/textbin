package main

import (
	"log"
	"net/http"
)

const port = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// start and listen to server
	log.Println("Starting server on", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
