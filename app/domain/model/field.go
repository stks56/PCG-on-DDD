package model

type Field struct {
	Deck        *Deck
	BattleField *BattleField
	Bench       *Bench
	Side        *Side
	Trash       *Trash
	Hand        *Hand
}
