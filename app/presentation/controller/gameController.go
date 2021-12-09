package controller

import (
	"net/http"
	"stks56/PCG-on-DDD/app/application/usecase"
	"stks56/PCG-on-DDD/app/domain/model"
	"stks56/PCG-on-DDD/app/infrastructure/repository"

	"github.com/labstack/echo/v4"
)

type GameControlller struct{}

func (gc GameControlller) Start() echo.HandlerFunc {
	return func(c echo.Context) error {
		// デッキ作成ができるまでプリセットを使う
		yourDeck := model.NewPresetDeck()
		opponentDeck := model.NewPresetDeck()

		gu := usecase.GameUsecase{GameModel: &model.Game{}}
		game := gu.StartGame(yourDeck, opponentDeck)

		gr := repository.NewGameRepository()
		if err := gr.Insert(game); err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed insert db")
		}

		return c.JSON(http.StatusOK, game.YourField.Hand)
	}
}
