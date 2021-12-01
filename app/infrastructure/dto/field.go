package dto

type Field struct {
	Id          uint `gorm:"primaryKey"`
	GameId      uint
	Deck        Deck
	BattleField BattleField
	Bench       Bench
	Side        Side
	Trash       Trash
	Hand        Hand
}

type Deck struct {
	Id        uint
	FieldId   uint
	GameCards []GameCard
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

type GameCard struct {
	DeckId uint
	CardId uint
}
