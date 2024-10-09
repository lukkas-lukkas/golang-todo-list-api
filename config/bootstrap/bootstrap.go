package bootstrap

import (
	"github.com/lukkas-lukkas/golang-todo-list-api/config/service_providers"
)

func Boot() {
	service_providers.NewEnvProvider().Boot()
	service_providers.NewDBProvider().Boot()
}
