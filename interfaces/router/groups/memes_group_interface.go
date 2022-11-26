package groups

import "github.com/labstack/echo/v4"

type MemeGroup interface {
	Resources(group *echo.Group)
}
