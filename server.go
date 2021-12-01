package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"stks56/PCG-on-DDD/app/infrastructure"
	"stks56/PCG-on-DDD/app/presentation/controller"
)

func main() {
	infrastructure.InitDB()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	info := e.Group("/info")
	{
		card := controller.CardController{}
		info.GET("/card/:id", card.Show())

		field := controller.FieldController{}
		info.GET("/field", field.Show())
	}

	game := e.Group("/game")
	{
		gameCtrl := controller.GameControlller{}
		game.POST("/start", gameCtrl.Start())
	}

	e.Logger.Fatal(e.Start(":1323"))
}
