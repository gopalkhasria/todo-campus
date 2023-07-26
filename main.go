package main

import (
	"net/http"
	"todo/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.HelloWord)
	http.ListenAndServe(":3000", r)
	
}
