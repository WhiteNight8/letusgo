package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hello world"))
}

func letusgoView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the query parameter "id"
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "letusgo view id: %d", id)
}

func letusgoCreate(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check if it is POST
	if r.Method != "POST" {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("letusgo create"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/letusgo/view", letusgoView)
	mux.HandleFunc("/letusgo/create", letusgoCreate)

	log.Print("server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
