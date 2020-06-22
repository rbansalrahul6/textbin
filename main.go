package main

import (
	"log"
	"net/http"
)

const port = ":4000"

// handler func for home
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from textbin"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// start and listen to server
	log.Println("Starting server on", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
