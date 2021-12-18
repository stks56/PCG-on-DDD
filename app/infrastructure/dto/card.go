package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Card struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Deck        []Deck
	BattleField []BattleField
	Bench       []Bench
	Side        []Side
	Trash       []Trash
	Hand        []Hand
}

type Cards []Card

func (card *Card) ConvertToModel() *model.Card {
	return &model.Card{
		Id:   int(card.Id),
		Name: card.Name,
	}
}
