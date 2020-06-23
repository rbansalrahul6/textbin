package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// add app dependencies here
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// command line flags for config
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// initialize loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize our app
	app := &application{
		infoLog:  infoLog,
		errorLog: errLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// start and listen to server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := server.ListenAndServe()
	errLog.Fatal(err)
}
