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
	yDeck := make([]dto.Deck, 0, 60)
	for _, c := range game.YourField.Deck.Cards {
		yDeck = append(yDeck, dto.Deck{CardId: uint(c.Id)})
	}
	oDeck := make([]dto.Deck, 0, 60)
	for _, c := range game.OpponentField.Deck.Cards {
		oDeck = append(oDeck, dto.Deck{CardId: uint(c.Id)})
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
				Deck: yDeck,
				Hand: yHand,
			},
			{
				Deck: oDeck,
				Hand: oHand,
			},
		},
		Status: uint8(game.Status),
	}

	tx := gr.db.Create(dto)
	return tx.Error
}

func (gr *gameRepository) Get() (*model.Game, error) {
	dto := &dto.Game{
		Fields: []dto.Field{
			{
				Deck:        []dto.Deck{},
				BattleField: dto.BattleField{},
				Bench:       []dto.Bench{},
				Side:        []dto.Side{},
				Trash:       []dto.Trash{},
				Hand:        []dto.Hand{},
			},
			{
				Deck:        []dto.Deck{},
				BattleField: dto.BattleField{},
				Bench:       []dto.Bench{},
				Side:        []dto.Side{},
				Trash:       []dto.Trash{},
				Hand:        []dto.Hand{},
			},
		},
	}
	gr.db.First(dto)
	return dto.ConvertToModel(), nil
}
