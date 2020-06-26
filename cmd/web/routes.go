package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippetForm)))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.showSnippet)))

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
