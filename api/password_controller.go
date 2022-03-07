package api

import (
	"net/http"

	"github.com/tmdgo/router"
	"github.com/toolfordev/local-api-encrypted-variables/application"
	"github.com/toolfordev/local-api-encrypted-variables/models"
)

type PasswordController struct {
}

func (PasswordController) GetRoutes() []router.Route {
	return []router.Route{
		{Path: "/password/set", Method: http.MethodPatch, UseRequestModel: true, RequestModel: func() interface{} { return &models.Password{} }, HandleFunc: func(service *application.PasswordService, model *models.Password) (result router.Result, err router.Error) {
			errFunc := service.Set(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.StatusCode = http.StatusOK
			return
		}},
		{Path: "/password/lock", Method: http.MethodPatch, UseRequestModel: false, HandleFunc: func(service *application.PasswordService) (result router.Result, err router.Error) {
			service.Lock()
			result.StatusCode = http.StatusOK
			return
		}},
		{Path: "/password/unlock", Method: http.MethodPatch, UseRequestModel: true, RequestModel: func() interface{} { return &models.Password{} }, HandleFunc: func(service *application.PasswordService, model *models.Password) (result router.Result, err router.Error) {
			errFunc := service.Unlock(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.StatusCode = http.StatusOK
			return
		}},
	}
}
