package usecase

import "stks56/PCG-on-DDD/app/domain/model"

type GameUsecase struct {
	GameModel *model.Game
}

func (gu *GameUsecase) StartGame(yourDeck, opponentDeck *model.Deck) *model.Game {
	gu.GameModel.YourField = &model.Field{}
	gu.GameModel.OpponentField = &model.Field{}
	gu.GameModel.YourField.Deck = yourDeck
	gu.GameModel.OpponentField.Deck = opponentDeck
	return gu.GameModel
}
