package application

import (
	"github.com/toolfordev/local-api-encrypted-variables/models"
	"github.com/toolfordev/local-api-encrypted-variables/persistence"
)

type EncryptedVariableService struct {
	repository      *persistence.EncryptedVariableRepository
	passwordService *PasswordService
}

func (service *EncryptedVariableService) Init(passwordService *PasswordService, repository *persistence.EncryptedVariableRepository) {
	service.repository = repository
	service.passwordService = passwordService
}

func (service *EncryptedVariableService) Create(model *models.EncryptedVariable) (err error) {
	err = service.passwordService.Encrypt(model)
	if err != nil {
		return
	}
	err = service.repository.Create(model)
	return
}

func (service *EncryptedVariableService) Update(model *models.EncryptedVariable) (err error) {
	err = service.repository.Update(model)
	return
}

func (service *EncryptedVariableService) GetByID(id uint) (model models.EncryptedVariable, err error) {
	model, err = service.repository.GetByID(id)
	if err != nil {
		return
	}
	err = service.passwordService.Decrypt(&model)
	return
}

func (service *EncryptedVariableService) GetAll() (models []models.EncryptedVariable, err error) {
	models, err = service.repository.GetAll()
	if err != nil {
		return
	}
	models, err = service.passwordService.DecryptAll(models)
	return
}

func (service *EncryptedVariableService) DeleteByID(id uint) (err error) {
	err = service.repository.DeleteByID(id)
	return
}
