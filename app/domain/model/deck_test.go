package model

import "testing"

func TestShuffle(t *testing.T) {
	deck := Deck{
		Cards: []Card{
			{Id: 1},
			{Id: 2},
			{Id: 3},
			{Id: 4},
			{Id: 5},
		},
	}
	testingSeed := int64(1638647744204023000)
	expected := 1
	actual := deck.Cards[0].Id
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	deck.Shuffle(testingSeed)
	actual = deck.Cards[0].Id

	if actual == expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
