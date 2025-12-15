package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	//define a new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/letusgo/view", letusgoView)
	mux.HandleFunc("/letusgo/create", letusgoCreate)

	log.Printf("server is running on port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
