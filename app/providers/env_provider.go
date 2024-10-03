package providers

import "github.com/joho/godotenv"

type EnvProvider struct{}

func NewEnvProvider() *EnvProvider {
	return &EnvProvider{}
}

func (p EnvProvider) Boot() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
