package main

import (
	"fmt"
	"github.com/lukkas-lukkas/golang-todo-list-api/bin"
	"github.com/lukkas-lukkas/golang-todo-list-api/routes"
	"os"
)

func main() {
	bin.Boot()

	fmt.Println("APP_ENV", os.Getenv("APP_ENV"))

	routes.ExecuteApi()
}
