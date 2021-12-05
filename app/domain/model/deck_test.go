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

func TestDrawCards(t *testing.T) {
	deck := Deck{
		Cards: []Card{
			{Id: 1},
			{Id: 2},
			{Id: 3},
			{Id: 4},
			{Id: 5},
		},
	}
	var actual int
	var expected int

	cards := deck.DrawCards(1)
	actual = cards[0].Id
	expected = 1
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
	actual = len(deck.Cards)
	expected = 4
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	cards = deck.DrawCards(3)
	actual = cards[0].Id
	expected = 2
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
	actual = len(deck.Cards)
	expected = 1
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	cards = deck.DrawCards(3)
	actual = len(cards)
	expected = 0
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
