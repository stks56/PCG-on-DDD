package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Side struct {
	FieldId uint
	CardId  uint
}

func (side *Side) ConvertToModel() *model.Card {
	return &model.Card{Id: int(side.CardId)}
}
