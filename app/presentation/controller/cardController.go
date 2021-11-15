package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CardController struct{}

func (cc CardController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
