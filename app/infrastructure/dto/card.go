package dto

type Card struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Deck        []Deck
	BattleField []BattleField
	Bench       []Bench
	Side        []Side
	Trash       []Trash
	Hand        []Hand
}

type Cards []Card
