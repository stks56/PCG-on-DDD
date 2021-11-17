package controller

import (
	"github.com/labstack/echo/v4"
)

type FieldController struct{}

func (fc FieldController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
