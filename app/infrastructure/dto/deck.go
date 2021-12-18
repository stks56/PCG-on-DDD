package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Deck struct {
	FieldId uint
	CardId  uint
}

func (deck *Deck) ConvertToModel() *model.Card {
	return &model.Card{Id: int(deck.CardId)}
}
