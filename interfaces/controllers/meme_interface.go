package controllers

import "github.com/labstack/echo/v4"

type MemeController interface {
	GenerateMeme(ctx echo.Context) error
}
