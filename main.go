package main

import (
	"net/http"
	"todo/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.ServeHtml)
	fs := http.FileServer(http.Dir("./static/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.ListenAndServe(":3000", r)
	
}
