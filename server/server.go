package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	echo := echo.New()

	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000"},
	}))
	echo.Use(middleware.Recover())

	Router(echo)

	echo.Logger.Fatal(echo.Start(":8000"))
}
