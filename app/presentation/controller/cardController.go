package controller

import (
	"net/http"
	"strconv"

	"stks56/PCG-on-DDD/app/infrastructure/repository"

	"github.com/labstack/echo/v4"
)

type CardController struct{}

func (cc CardController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		cr := repository.NewCardRepository()
		id, _ := strconv.Atoi(c.Param("id"))

		card, err := cr.GetByID(id)
		if err != nil {
			panic("error")
		}

		return c.JSON(http.StatusOK, card)
	}
}
