package controller

import (
	"net/http"
	"path"
	"text/template"
)

func HelloWord(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello word"))
}

func ServeHtml(w http.ResponseWriter, request *http.Request) {
	var filepath = path.Join("page", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
