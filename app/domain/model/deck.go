package model

type Deck struct {
	Cards *[60]Card
}

func NewPresetDeck() *Deck {
	var cards [60]Card
	for i := 0; i < 60; i++ {
		cards[i] = Card{Id: i % 3, Name: ""}
	}
	return &Deck{
		Cards: (*[60]Card)(&cards),
	}
}
