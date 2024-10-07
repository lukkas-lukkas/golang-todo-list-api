package config

import "github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/db"

func Load() {
	EnvConfig()
	db.InitDB()
}
