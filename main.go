package main

import (
	"github.com/tmdgo/dependencies"
	"github.com/toolfordev/local-api-encrypted-variables/api"
	"github.com/toolfordev/local-api-encrypted-variables/application"
	"github.com/toolfordev/local-api-encrypted-variables/persistence"
)

func main() {
	manager := &dependencies.Manager{}
	manager.Init()
	manager.Add(&persistence.ToolForDevDatabase{})
	manager.Add(&persistence.PasswordHashRepository{})
	manager.Add(&persistence.EncryptedVariableRepository{})
	manager.Add(&application.PasswordService{})
	manager.Add(&application.EncryptedVariableService{})

	manager.CallFunc(func(database *persistence.ToolForDevDatabase) {

	})

	api := api.EncryptedVariableApi{}
	api.Init(manager)
}
