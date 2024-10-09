package mysql

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"log"
	"os"
)

var Connection *sql.DB

func InitConnection() {
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

	Connection = db
}

func CloseConnection() {
	err := Connection.Close()
	if err != nil {
		log.Fatalln("DB CLOSE ERROR", err)
	}
}
