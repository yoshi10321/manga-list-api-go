package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/manga-list-api-go/handler"
	"github.com/manga-list-api-go/model"
)

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	mangaHandler := handler.NewMangaHandler(model.NewMangaRepo())
	e.GET("/mangas", mangaHandler.GetMangaList)
	e.POST("/mangas", mangaHandler.AddMangaItem)
	e.POST("/mangas/:id", mangaHandler.UpdateMangaItem)
}
