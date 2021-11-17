package repository

import (
	"stks56/PCG-on-DDD/app/domain/model"
	"stks56/PCG-on-DDD/app/domain/repository"
	"stks56/PCG-on-DDD/app/infrastructure"
	"stks56/PCG-on-DDD/app/infrastructure/dto"

	"gorm.io/gorm"
)

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository() repository.CardRepository {
	return &cardRepository{
		db: infrastructure.GetDB(),
	}
}

func (cr *cardRepository) GetByID(id int) (*model.Card, error) {
	dto := dto.Card{}
	cr.db.Where("id = ?", id).First(&dto)

	card := model.Card{
		Id:   id,
		Name: dto.Name,
	}
	return &card, nil
}
