package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"letgo.snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// use the snippets model to get the 10 most recently created snippets
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v", snippet)
	}

	//intialize a slice containing the paths to the two files
	// 	files := []string{
	// 		"./ui/html/pages/home.html",
	// 		"./ui/html/pages/base.html",
	// 		"./ui/html/partials/nav.html",
	// 	}

	// 	// read the template file
	// 	ts, err := template.ParseFiles(files...)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}

	// 	//write the template content as the response body
	// 	err = ts.ExecuteTemplate(w, "base", nil)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 	}
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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
