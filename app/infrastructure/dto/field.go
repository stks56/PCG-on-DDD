package dto

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

type Deck struct {
	FieldId uint
	CardId  uint
}

type BattleField struct {
	FieldId uint
	CardId  uint
}

type Bench struct {
	FieldId uint
	CardId  uint
}

type Side struct {
	FieldId uint
	CardId  uint
}

type Trash struct {
	FieldId uint
	CardId  uint
}

type Hand struct {
	FieldId uint
	CardId  uint
}
