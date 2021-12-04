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
}
