package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("CustomMiddleware . . .")
	return func(c echo.Context) error {
		return next(c)
	}
}
