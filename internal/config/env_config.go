package config

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err)
	}
}
