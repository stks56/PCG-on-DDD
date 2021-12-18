package model

type Bench struct {
	Cards []Card
}

func (banch *Bench) Set(card *Card) {
	banch.Cards = append(banch.Cards, *card)
}
