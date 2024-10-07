package controllers

import "net/http"

func RegisterController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register user!"))
}
