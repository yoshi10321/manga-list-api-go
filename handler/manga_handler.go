package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/manga-list-api-go/model"
	req "github.com/manga-list-api-go/request"
)

type MangaHandler struct {
	mangaRepo model.MangaRepoInterface
}

func NewMangaHandler(mangaRepo model.MangaRepoInterface) MangaHandler {
	return MangaHandler{
		mangaRepo: mangaRepo,
	}
}

func (handler MangaHandler) GetMangaList(c echo.Context) (err error) {
	mangaList, err := handler.mangaRepo.FindAll()
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	return c.JSON(http.StatusOK, mangaList)
}

func (handler MangaHandler) AddMangaItem(c echo.Context) (err error) {
	request := new(req.MangaCreate)
	if err = c.Bind(request); err != nil {
		log.Printf("error: %v", err)
		return
	}
	err = handler.mangaRepo.Insert(request.Title, request.ImgURL)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	return nil
}

func (handler MangaHandler) UpdateMangaItem(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	request := new(req.MangaUpdate)
	if err = c.Bind(request); err != nil {
		log.Printf("error: %v", err)
		return
	}
	manga, err := handler.mangaRepo.Find(id)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	if request.Title != "" {
		manga.Title = request.Title
	}

	if request.ImgURL != "" {
		manga.ImgURL = request.ImgURL
	}

	if request.Number != 0 {
		manga.Number = request.Number
	}

	handler.mangaRepo.Update(manga)

	return nil
}
