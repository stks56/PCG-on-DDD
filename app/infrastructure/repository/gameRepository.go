package repository

import (
	"stks56/PCG-on-DDD/app/domain/model"
	"stks56/PCG-on-DDD/app/domain/repository"
	"stks56/PCG-on-DDD/app/infrastructure"
	"stks56/PCG-on-DDD/app/infrastructure/dto"

	"gorm.io/gorm"
)

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository() repository.GameRepository {
	return &gameRepository{
		db: infrastructure.GetDB(),
	}
}

func (gr *gameRepository) Insert(game *model.Game) error {
	yCards := make([]dto.GameCard, 0, 60)
	for _, c := range game.YourField.Deck.Cards {
		yCards = append(yCards, dto.GameCard{CardId: uint(c.Id)})
	}
	oCards := make([]dto.GameCard, 0, 60)
	for _, c := range game.OpponentField.Deck.Cards {
		oCards = append(oCards, dto.GameCard{CardId: uint(c.Id)})
	}
	yHand := make([]dto.Hand, 0, 7)
	for _, c := range game.YourField.Hand.Cards {
		yHand = append(yHand, dto.Hand{CardId: uint(c.Id)})
	}
	oHand := make([]dto.Hand, 0, 7)
	for _, c := range game.OpponentField.Hand.Cards {
		oHand = append(oHand, dto.Hand{CardId: uint(c.Id)})
	}

	dto := &dto.Game{
		Fields: []dto.Field{
			{
				Deck: dto.Deck{GameCards: yCards},
				Hand: yHand,
			},
			{
				Deck: dto.Deck{GameCards: oCards},
				Hand: oHand,
			},
		},
	}

	tx := gr.db.Create(dto)
	return tx.Error
}
