package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Field struct {
	Id          uint `gorm:"primaryKey"`
	GameId      uint
	Deck        []Deck
	BattleField BattleField
	Bench       []Bench
	Side        []Side
	Trash       []Trash
	Hand        []Hand
}

func (field *Field) ConvertToModel() *model.Field {
	deckModel := &model.Deck{}
	for _, deck := range field.Deck {
		deckModel.Cards = append(deckModel.Cards, *deck.ConvertToModel())
	}
	benchModel := &model.Bench{}
	for _, bench := range field.Bench {
		benchModel.Cards = append(benchModel.Cards, *bench.ConvertToModel())
	}
	sideModel := &model.Side{}
	for _, side := range field.Side {
		sideModel.Cards = append(sideModel.Cards, *side.ConvertToModel())
	}
	trashModel := &model.Trash{}
	for _, trash := range field.Trash {
		trashModel.Cards = append(trashModel.Cards, *trash.ConvertToModel())
	}
	handModel := &model.Hand{}
	for _, hand := range field.Hand {
		handModel.Cards = append(handModel.Cards, *hand.ConvertToModel())
	}

	return &model.Field{
		Deck:        deckModel,
		BattleField: field.BattleField.ConvertToModel(),
		Bench:       benchModel,
		Side:        sideModel,
		Trash:       trashModel,
		Hand:        handModel,
	}
}
