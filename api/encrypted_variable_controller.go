package api

import (
	"net/http"
	"strconv"

	"github.com/tmdgo/router"
	"github.com/toolfordev/local-api-encrypted-variables/application"
	"github.com/toolfordev/local-api-encrypted-variables/models"
)

type EncryptedVariableController struct {
}

func (EncryptedVariableController) GetRoutes() []router.Route {
	return []router.Route{
		{Path: "/encrypted-variable", Method: http.MethodPost, UseRequestModel: true, RequestModel: func() interface{} { return &models.EncryptedVariable{} }, HandleFunc: func(service *application.EncryptedVariableService, model *models.EncryptedVariable) (result router.Result, err router.Error) {
			errFunc := service.Create(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.Model = model
			result.StatusCode = http.StatusCreated
			return
		}},
		{Path: "/encrypted-variable", Method: http.MethodPut, UseRequestModel: true, RequestModel: func() interface{} { return &models.EncryptedVariable{} }, HandleFunc: func(service *application.EncryptedVariableService, model *models.EncryptedVariable) (result router.Result, err router.Error) {
			errFunc := service.Update(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.Model = model
			result.StatusCode = http.StatusOK
			return
		}},
		{Path: "/encrypted-variable", Method: http.MethodGet, HandleFunc: func(service *application.EncryptedVariableService) (result router.Result, err router.Error) {
			models, errFunc := service.GetAll()
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			result.Model = models
			return
		}},
		{Path: "/encrypted-variable/{id:[0-9]+}", Method: http.MethodGet, UseVars: true, HandleFunc: func(service *application.EncryptedVariableService, vars *router.Vars) (result router.Result, err router.Error) {
			id, errFunc := strconv.ParseUint(vars.Value["id"], 10, 32)
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			model, errFunc := service.GetByID(uint(id))
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			result.Model = model
			return
		}},
		{Path: "/encrypted-variable/{id:[0-9]+}", Method: http.MethodDelete, UseVars: true, HandleFunc: func(service *application.EncryptedVariableService, vars *router.Vars) (result router.Result, err router.Error) {
			id, errFunc := strconv.ParseUint(vars.Value["id"], 10, 32)
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			errFunc = service.DeleteByID(uint(id))
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			return
		}},
	}
}
