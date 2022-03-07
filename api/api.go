package api

import (
	"github.com/tmdgo/dependencies"
	"github.com/tmdgo/router"
)

type EncryptedVariableApi struct {
	router router.Router
}

func (api *EncryptedVariableApi) Init(manager *dependencies.Manager) {
	api.router.Init(manager)
	api.router.AddController(EncryptedVariableController{})
	api.router.ListenAndServe()
}
