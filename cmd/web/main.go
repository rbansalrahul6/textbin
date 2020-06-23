package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// command line flags for config
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// start and listen to server
	log.Println("Starting server on", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
