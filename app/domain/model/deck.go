package model

import (
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func NewPresetDeck() *Deck {
	cards := make([]Card, 0, 60)
	for i := 0; i < 60; i++ {
		cards = append(cards, Card{Id: i%3 + 1, Name: ""})
	}
	return &Deck{
		Cards: cards,
	}
}

func (deck *Deck) Shuffle(seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

func (deck *Deck) DrawCards(sheets int) []Card {
	cards := make([]Card, 0, sheets)

	// If try to draw more cards than the sheets in deck. game will end.
	// but this implement as can't draw cards.
	if len(deck.Cards) < sheets {
		return cards
	}

	for i := 0; i < sheets; i++ {
		var card Card
		card, deck.Cards = deck.Cards[0], deck.Cards[1:]
		cards = append(cards, card)
	}
	return cards
}
