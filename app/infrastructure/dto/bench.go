package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Bench struct {
	FieldId uint
	CardId  uint
}

func (bench *Bench) ConvertToModel() *model.Card {
	return &model.Card{Id: int(bench.CardId)}
}
