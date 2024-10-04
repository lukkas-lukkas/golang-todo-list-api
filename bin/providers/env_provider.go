package bin

import (
	"github.com/joho/godotenv"
	"log"
)

type EnvProvider struct{}

func NewEnvProvider() *EnvProvider {
	return &EnvProvider{}
}

func (p EnvProvider) Boot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err)
	}
}
