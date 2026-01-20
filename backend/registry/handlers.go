package registry

import "github.com/labstack/echo/v4"

type applicantHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
}
