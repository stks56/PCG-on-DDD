package model

type Card struct {
	Id   int
	Name string
}

func NewCard(id int, name string) *Card {
	return &Card{
		Id:   id,
		Name: name,
	}
}
