package persistence

import (
	"github.com/toolfordev/local-api-encrypted-variables/models"
	"gorm.io/gorm"
)

type PasswordHashEntity struct {
	gorm.Model
	Value string
}

type PasswordHashConverter struct {
}

func (PasswordHashConverter) ToEntity(model models.PasswordHash) (entity PasswordHashEntity) {
	entity.Value = model.Value
	return
}

func (PasswordHashConverter) ToModel(entity PasswordHashEntity) (model models.PasswordHash) {
	model.Value = entity.Value
	return
}

func (converter PasswordHashConverter) ToModels(entities []PasswordHashEntity) (modelsResult []models.PasswordHash) {
	modelsResult = make([]models.PasswordHash, len(entities))
	for index, entity := range entities {
		model := converter.ToModel(entity)
		modelsResult[index] = model
	}
	return
}

type PasswordHashRepository struct {
	database  *ToolForDevDatabase
	converter *PasswordHashConverter
}

func (repository *PasswordHashRepository) Init(database *ToolForDevDatabase) {
	repository.database = database
	repository.database.RegisterEntities(PasswordHashEntity{})

	repository.converter = &PasswordHashConverter{}
}

func (repository *PasswordHashRepository) Create(model *models.PasswordHash) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Create(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *PasswordHashRepository) Update(model *models.PasswordHash) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Update(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *PasswordHashRepository) DeleteByID(id uint) (err error) {
	err = repository.database.DeleteByID(&PasswordHashEntity{}, id)
	return
}

func (repository *PasswordHashRepository) GetByID(id uint) (model models.PasswordHash, err error) {
	entity := PasswordHashEntity{}
	err = repository.database.SelectByID(&entity, id)
	if err != nil {
		return
	}
	model = repository.converter.ToModel(entity)
	return
}

func (repository *PasswordHashRepository) GetAll() (models []models.PasswordHash, err error) {
	entities := make([]PasswordHashEntity, 0)
	err = repository.database.SelectAll(&entities)
	if err != nil {
		return
	}
	models = repository.converter.ToModels(entities)
	return
}
