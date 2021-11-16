package repository

import (
	"stks56/PCG-onDDD/app/domain/model"
	"stks56/PCG-onDDD/app/domain/repository"
)

type cardRepository struct{}

func NewCardRepository() repository.CardRepository {
	return &cardRepository{}
}

func (c *cardRepository) GetByID(id int) (*model.Card, error) {
	// tmp
	card := model.NewCard(id, "foo")
	return card, nil
}