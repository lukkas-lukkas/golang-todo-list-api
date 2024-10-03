package controllers

import "net/http"

func HelloWorldController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
