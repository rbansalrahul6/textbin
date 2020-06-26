package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rbansalrahul6/textbin/pkg/forms"
	"github.com/rbansalrahul6/textbin/pkg/models"
)

// handler func for home
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// template rendering
	data := &templateData{Snippets: s}
	app.render(w, r, "home.page.tmpl", data)
}

// handler func to show a recorded snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	// Read from DB
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	// render template
	data := &templateData{Snippet: s}
	app.render(w, r, "snippet.page.tmpl", data)
}

// handler func to crate a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }
	// w.Write([]byte("Create a new snippet..."))
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// title := r.PostForm.Get("title")
	// content := r.PostForm.Get("content")
	// expires := r.PostForm.Get("expires")
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	// // user input validation
	// errors := make(map[string]string)
	// // check non-null & max len constraint
	// if strings.TrimSpace(title) == "" {
	// 	errors["title"] = "This field cannot be blank"
	// } else if utf8.RuneCountInString(title) > 100 {
	// 	errors["title"] = "This field is too long (maximum is 100 characters)"
	// }

	// // validations for content field
	// if strings.TrimSpace(content) == "" {
	// 	errors["content"] = "This field cannot be blank"
	// }

	// // validations for expires field
	// if strings.TrimSpace(expires) == "" {
	// 	errors["expires"] = "This field cannot be blank"
	// } else if expires != "1" && expires != "7" && expires != "365" {
	// 	errors["expires"] = "This field is invalid"
	// }

	if !form.Valid() {
		// TODO: re-display form with errors
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		// fmt.Fprint(w, errors)
		return
	}

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

// handler func for snippet form
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
