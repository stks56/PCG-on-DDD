package model

import (
	"testing"
)

func TestAddCards(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		},
	}
	cards := []Card{
		{Id: 4},
		{Id: 5},
	}
	hand.AddCards(cards)

	actual := len(hand.Cards)
	expected := 5
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
