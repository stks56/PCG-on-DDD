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
		cardCtrl := controller.CardController{}
		info.GET("/card/:id", cardCtrl.Show())
	}

	game := e.Group("/game")
	{
		gameCtrl := controller.GameControlller{}
		game.POST("/start", gameCtrl.Start())

		fieldCtrl := controller.FieldController{}
		game.GET("/field", fieldCtrl.Show())
	}

	e.Logger.Fatal(e.Start(":1323"))
}
