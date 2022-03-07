package persistence

import (
	"github.com/toolfordev/local-api-encrypted-variables/models"
	"gorm.io/gorm"
)

type EncryptedVariableEntity struct {
	gorm.Model
	Name           string `gorm:"unique"`
	EncryptedValue string
}

type EncryptedVariableConverter struct {
}

func (EncryptedVariableConverter) ToEntity(model models.EncryptedVariable) (entity EncryptedVariableEntity) {
	entity.ID = model.ID
	entity.Name = model.Name
	entity.EncryptedValue = model.EncryptedValue
	return
}

func (EncryptedVariableConverter) ToModel(entity EncryptedVariableEntity) (model models.EncryptedVariable) {
	model.ID = entity.ID
	model.Name = entity.Name
	model.EncryptedValue = entity.EncryptedValue
	return
}

func (converter EncryptedVariableConverter) ToModels(entities []EncryptedVariableEntity) (modelsResult []models.EncryptedVariable) {
	modelsResult = make([]models.EncryptedVariable, len(entities))
	for index, entity := range entities {
		model := converter.ToModel(entity)
		modelsResult[index] = model
	}
	return
}

type EncryptedVariableRepository struct {
	database  *ToolForDevDatabase
	converter *EncryptedVariableConverter
}

func (repository *EncryptedVariableRepository) Init(database *ToolForDevDatabase) {
	repository.database = database
	repository.database.RegisterEntities(EncryptedVariableEntity{})

	repository.converter = &EncryptedVariableConverter{}
}

func (repository *EncryptedVariableRepository) Create(model *models.EncryptedVariable) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Create(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *EncryptedVariableRepository) Update(model *models.EncryptedVariable) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Update(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *EncryptedVariableRepository) DeleteByID(id uint) (err error) {
	err = repository.database.DeleteByID(&EncryptedVariableEntity{}, id)
	return
}

func (repository *EncryptedVariableRepository) GetByID(id uint) (model models.EncryptedVariable, err error) {
	entity := EncryptedVariableEntity{}
	err = repository.database.SelectByID(&entity, id)
	if err != nil {
		return
	}
	model = repository.converter.ToModel(entity)
	return
}

func (repository *EncryptedVariableRepository) GetAll() (models []models.EncryptedVariable, err error) {
	entities := make([]EncryptedVariableEntity, 0)
	err = repository.database.SelectAll(&entities)
	if err != nil {
		return
	}
	models = repository.converter.ToModels(entities)
	return
}
