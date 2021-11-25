package controller

import (
	"net/http"
	"stks56/PCG-on-DDD/app/application/usecase"
	"stks56/PCG-on-DDD/app/domain/model"

	"github.com/labstack/echo/v4"
)

type GameControlller struct{}

func (cc GameControlller) Start() echo.HandlerFunc {
	return func(c echo.Context) error {
		// デッキ作成ができるまでプリセットを使う
		yourDeck := model.NewPresetDeck()
		opponentDeck := model.NewPresetDeck()

		gu := usecase.GameUsecase{GameModel: &model.Game{}}
		game := gu.StartGame(yourDeck, opponentDeck)

		return c.JSON(http.StatusOK, nil)
	}
}
