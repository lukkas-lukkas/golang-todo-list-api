package main

import (
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/persistence/mysql"
	"github.com/lukkas-lukkas/golang-todo-list-api/config/bootstrap"
	"github.com/lukkas-lukkas/golang-todo-list-api/config/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	bootstrap.Boot()
	defer mysql.CloseConnection()

	log.Fatal(http.ListenAndServe(
		":"+os.Getenv("API_SERVER_PORT"),
		routes.CreateRouter(),
	))
}
