package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GameControlller struct{}

func (cc GameControlller) Start() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
