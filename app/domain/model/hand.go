package model

type Hand struct {
	Cards []Card
}

func (hand *Hand) AddCards(cards []Card) {
	hand.Cards = append(hand.Cards, cards...)
}
