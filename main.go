package main

import (
	"log"
	"net/http"
)

const port = ":4000"

// handler func for home
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from textbin"))
}

// handler func to show a recorded snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a particualr snippet..."))
}

// handler func to crate a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// start and listen to server
	log.Println("Starting server on", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
