package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Post("/snippet/create", app.session.Enable(app.requireAuth(http.HandlerFunc(app.createSnippet))))
	mux.Get("/snippet/create", app.session.Enable(app.requireAuth(http.HandlerFunc(app.createSnippetForm))))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.showSnippet)))

	mux.Get("/user/signup", app.session.Enable(http.HandlerFunc(app.singupUserForm)))
	mux.Post("/user/signup", app.session.Enable(http.HandlerFunc(app.signupUser)))
	mux.Get("/user/login", app.session.Enable(http.HandlerFunc(app.loginUserForm)))
	mux.Post("/user/login", app.session.Enable(http.HandlerFunc(app.loginUser)))
	mux.Post("/user/logout", app.session.Enable(app.requireAuth(http.HandlerFunc(app.logoutUser))))

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
