package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manga-list-api-go/handler"
	"github.com/manga-list-api-go/model"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000"},
	}))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	mangaHandler := handler.NewMangaHandler(model.NewMangaRepo())
	e.GET("/mangas", mangaHandler.GetMangaList)
	e.POST("/mangas", mangaHandler.AddMangaItem)
	e.POST("/mangas/:id", mangaHandler.UpdateMangaItem)

	e.Logger.Fatal(e.Start(":8000"))
}
