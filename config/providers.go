package config

import "github.com/lukkas-lukkas/golang-todo-list-api/app/providers"

type AppProvider struct {
}

func NewAppProvider() *AppProvider {
	return &AppProvider{}
}

func (p AppProvider) Boot() {
	providers.NewEnvProvider().Boot()
}
