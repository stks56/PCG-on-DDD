package dto

import "stks56/PCG-on-DDD/app/domain/model"

type BattleField struct {
	FieldId uint
	CardId  uint
}

func (battleField *BattleField) ConvertToModel() *model.BattleField {
	return &model.BattleField{
		Card: &model.Card{Id: int(battleField.CardId)},
	}
}
