package controller

import "net/http"

func HelloWord(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello word"))
}
