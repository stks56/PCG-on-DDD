package repository

import (
	"stks56/PCG-on-DDD/app/domain/model"
	"stks56/PCG-on-DDD/app/domain/repository"
	"stks56/PCG-on-DDD/app/infrastructure"

	"gorm.io/gorm"
)

type fieldRepository struct {
	db *gorm.DB
}

func NewFieldRepository() repository.FieldRepository {
	return &fieldRepository{
		db: infrastructure.GetDB(),
	}
}

func (fr *fieldRepository) Get() (*model.Field, error) {}
