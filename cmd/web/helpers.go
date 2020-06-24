package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// method to handle server errors
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// method to handle client request errors
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// method to send 404 Not found error
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// method to render a template
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// get template from cache
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	// execute template
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
