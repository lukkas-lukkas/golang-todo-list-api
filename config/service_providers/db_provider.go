package service_providers

import "github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/persistence/mysql"

type DBProvider struct{}

func NewDBProvider() *DBProvider {
	return &DBProvider{}
}

func (p *DBProvider) Boot() {
	mysql.InitConnection()
}
