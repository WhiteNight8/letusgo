package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hello world"))
}

func letusgoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("letusgo view"))
}

func letusgoCreate(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check if it is POST
	if r.Method != "POST" {

		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
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
