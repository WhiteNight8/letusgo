package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// 使用相对于项目根目录的路径
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/letusgo/view", app.letusgoView)
	mux.HandleFunc("/letusgo/create", app.letusgoCreate)

	return mux
}
