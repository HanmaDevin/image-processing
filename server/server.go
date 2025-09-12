package server

import "github.com/labstack/echo/v4"

func NewServer() *echo.Echo {
	e := echo.New()

	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8000"))
}
