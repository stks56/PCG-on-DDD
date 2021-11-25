package dto

type Field struct {
	Id          uint `gorm:"primaryKey"`
	Deck        Deck
	BattleField BattleField
	Bench       Bench
	Side        Side
	Trash       Trash
	Hand        Hand
}

type Deck struct {
	FieldId uint
}

type BattleField struct {
	FieldId uint
}

type Bench struct {
	FieldId uint
}

type Side struct {
	FieldId uint
}

type Trash struct {
	FieldId uint
}

type Hand struct {
	FieldId uint
}
