package main

import (
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/db"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {
	config.Load()
	defer db.CloseDB()

	log.Fatal(http.ListenAndServe(
		":"+os.Getenv("API_SERVER_PORT"),
		config.CreateRouter(),
	))
}
