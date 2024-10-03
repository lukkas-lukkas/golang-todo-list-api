package bin

import bin "github.com/lukkas-lukkas/golang-todo-list-api/bin/providers"

func Boot() {
	bin.NewEnvProvider().Boot()
}
