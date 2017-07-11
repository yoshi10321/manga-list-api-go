package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/manga-list-api-go/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	helloHandler := handler.NewHelloHandler()
	e.GET("/hello", helloHandler.Hello)

	e.Logger.Fatal(e.Start(":8000"))
}
