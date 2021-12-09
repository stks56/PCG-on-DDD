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

	yourDeckCards := []model.Card{}
	for _, card := range dto.Fields[0].Deck {
		yourDeckCards = append(yourDeckCards, model.Card{Id: int(card.CardId)})
	}
	yourBattleFieldCard := &model.Card{Id: int(dto.Fields[0].BattleField.CardId)}
	yourBenchCards := []model.Card{}
	for _, card := range dto.Fields[0].Bench {
		yourBenchCards = append(yourBenchCards, model.Card{Id: int(card.CardId)})
	}
	yourSideCards := []model.Card{}
	for _, card := range dto.Fields[0].Side {
		yourSideCards = append(yourSideCards, model.Card{Id: int(card.CardId)})
	}
	yourTrashCards := []model.Card{}
	for _, card := range dto.Fields[0].Trash {
		yourTrashCards = append(yourTrashCards, model.Card{Id: int(card.CardId)})
	}
	yourHandCards := []model.Card{}
	for _, card := range dto.Fields[0].Hand {
		yourHandCards = append(yourHandCards, model.Card{Id: int(card.CardId)})
	}

	opponentDeckCards := []model.Card{}
	for _, card := range dto.Fields[1].Deck {
		opponentDeckCards = append(opponentDeckCards, model.Card{Id: int(card.CardId)})
	}
	opponentBattleFieldCard := &model.Card{Id: int(dto.Fields[1].BattleField.CardId)}
	opponentBenchCards := []model.Card{}
	for _, card := range dto.Fields[1].Bench {
		opponentBenchCards = append(opponentBenchCards, model.Card{Id: int(card.CardId)})
	}
	opponentSideCards := []model.Card{}
	for _, card := range dto.Fields[1].Side {
		opponentSideCards = append(opponentSideCards, model.Card{Id: int(card.CardId)})
	}
	opponentTrashCards := []model.Card{}
	for _, card := range dto.Fields[1].Trash {
		opponentTrashCards = append(opponentTrashCards, model.Card{Id: int(card.CardId)})
	}
	opponentHandCards := []model.Card{}
	for _, card := range dto.Fields[1].Hand {
		opponentHandCards = append(opponentHandCards, model.Card{Id: int(card.CardId)})
	}

	return &model.Game{
		YourField: &model.Field{
			Deck: &model.Deck{
				Cards: yourDeckCards,
			},
			BattleField: &model.BattleField{
				Card: yourBattleFieldCard,
			},
			Bench: &model.Bench{
				Cards: yourBenchCards,
			},
			Side: &model.Side{
				Cards: yourSideCards,
			},
			Trash: &model.Trash{
				Cards: yourTrashCards,
			},
			Hand: &model.Hand{
				Cards: yourHandCards,
			},
		},
		OpponentField: &model.Field{
			Deck: &model.Deck{
				Cards: opponentDeckCards,
			},
			BattleField: &model.BattleField{
				Card: opponentBattleFieldCard,
			},
			Bench: &model.Bench{
				Cards: opponentBenchCards,
			},
			Side: &model.Side{
				Cards: opponentSideCards,
			},
			Trash: &model.Trash{
				Cards: opponentTrashCards,
			},
			Hand: &model.Hand{
				Cards: opponentHandCards,
			},
		},
	}, nil
}
