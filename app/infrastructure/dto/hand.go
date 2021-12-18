package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Hand struct {
	FieldId uint
	CardId  uint
}

func (hand *Hand) ConvertToModel() *model.Card {
	return &model.Card{Id: int(hand.CardId)}
}
