package model

type BattleField struct {
	Card    *Card
	Reverse bool
}

func (battleField *BattleField) SetReverse(card *Card) {
	battleField.Card = card
	battleField.Reverse = true
}
