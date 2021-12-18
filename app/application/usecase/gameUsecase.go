package usecase

import (
	"stks56/PCG-on-DDD/app/domain/model"
	"time"
)

type GameUsecase struct {
	GameModel *model.Game
}

func (gu *GameUsecase) StartGame(yourDeck, opponentDeck *model.Deck) *model.Game {
	gu.GameModel.YourField = &model.Field{}
	gu.GameModel.OpponentField = &model.Field{}
	gu.GameModel.YourField.Deck = yourDeck
	gu.GameModel.OpponentField.Deck = opponentDeck

	gu.GameModel.YourField.Deck.Shuffle(time.Now().Unix())
	gu.GameModel.OpponentField.Deck.Shuffle(time.Now().Unix())

	gu.GameModel.YourField.Hand = &model.Hand{}
	gu.GameModel.OpponentField.Hand = &model.Hand{}

	CardsInYourHand := gu.GameModel.YourField.Deck.DrawCards(7)
	gu.GameModel.YourField.Hand.AddCards(CardsInYourHand)
	CardsInOpponentHand := gu.GameModel.OpponentField.Deck.DrawCards(7)
	gu.GameModel.OpponentField.Hand.AddCards(CardsInOpponentHand)

	return gu.GameModel
}

func (gu *GameUsecase) InitPokemon(battleFieldPokemonId int, benchesPokemonId []int) *model.Game {
	gu.GameModel.YourField.BattleField.SetReverse(&model.Card{Id: battleFieldPokemonId})
	return gu.GameModel
}
