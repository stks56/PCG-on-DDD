package controller

import (
	"net/http"
	"stks56/PCG-on-DDD/app/infrastructure/repository"

	"github.com/labstack/echo/v4"
)

type FieldController struct{}

func (fc FieldController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		fr := repository.NewFieldRepository()
		field, err := fr.Get()
		if err != nil {
			c.JSON(http.StatusNotFound, "Not Found Record")
		}

		return c.JSON(http.StatusOK, field)
	}
}
