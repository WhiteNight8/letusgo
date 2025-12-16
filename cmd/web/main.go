package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	//define a new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/letusgo/view", letusgoView)
	mux.HandleFunc("/letusgo/create", letusgoCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("server is running on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
