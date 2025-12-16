package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//intialize a slice containing the paths to the two files
	files := []string{
		"../../ui/html/pages/home.html",
		"../../ui/html/pages/base.html",
		"../../ui/html/partials/nav.html",
	}

	// read the template file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//write the template content as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (app *application) letusgoView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the query parameter "id"
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "letusgo view id: %d", id)
}

func (app *application) letusgoCreate(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check if it is POST
	if r.Method != "POST" {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("letusgo create"))
}
