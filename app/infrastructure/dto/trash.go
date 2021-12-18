package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Trash struct {
	FieldId uint
	CardId  uint
}

func (trash *Trash) ConvertToModel() *model.Card {
	return &model.Card{Id: int(trash.CardId)}
}
