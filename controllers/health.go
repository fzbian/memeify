package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Ok!")
}
