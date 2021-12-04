package model

import (
	"math/rand"
)

type Deck struct {
	Cards *[60]Card
}

func NewPresetDeck() *Deck {
	var cards [60]Card
	for i := 0; i < 60; i++ {
		cards[i] = Card{Id: i%3 + 1, Name: ""}
	}
	return &Deck{
		Cards: (*[60]Card)(&cards),
	}
}

func (deck *Deck) Shuffle(seed int64) {
	var shuffledCards [60]Card
	rand.Seed(seed)
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		shuffledCards[i], shuffledCards[j] = deck.Cards[j], deck.Cards[i]
	})
	deck.Cards = &shuffledCards
}
