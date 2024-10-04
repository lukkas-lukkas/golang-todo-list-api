package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lukkas-lukkas/golang-todo-list-api/bin"
	"github.com/lukkas-lukkas/golang-todo-list-api/routes"
	"log"
	"os"
)

func main() {
	bin.Boot()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalln("Error to open database connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Error pinging database", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln("DB CLOSE ERROR", err)
		}
	}(db)

	fmt.Println("Database connection successful")

	routes.ExecuteApi()
}
