package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/letusgo/view", letusgoView)
	mux.HandleFunc("/letusgo/create", letusgoCreate)

	log.Print("server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
