package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"letgo.snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	//create some variables holding dummy data
	//we will remove these later on
	title := "Home"
	content := "Hello, World!"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/letusgo/view?id=%d", id), http.StatusSeeOther)

	//intialize a slice containing the paths to the two files
	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/pages/base.html",
		"./ui/html/partials/nav.html",
	}

	// read the template file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//write the template content as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) letusgoView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the query parameter "id"
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// use the snippets model to get the data for a specific record based on its ID
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) letusgoCreate(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check if it is POST
	if r.Method != "POST" {

		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("letusgo create"))
}
