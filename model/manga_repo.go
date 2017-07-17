package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// MangaRepoInterface define what MangaRepo struct have some methods.
type MangaRepoInterface interface {
	Insert(title string, imgURL string) error
	Find(id int) (Manga, error)
	FindAll() ([]Manga, error)
	Update(Manga) error
}

type MangaRepo struct {
}

func NewMangaRepo() MangaRepoInterface {
	return MangaRepo{}
}

func (repo MangaRepo) Insert(title string, imgURL string) error {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		e := echo.New()
		e.Logger.Fatal("failed to connect database")
	}
	defer db.Close()

	db.Create(&Manga{Title: title, ImgURL: imgURL, Number: 1})

	return nil
}

func (repo MangaRepo) Find(id int) (Manga, error) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		e := echo.New()
		e.Logger.Fatal("failed to connect database")
	}
	defer db.Close()

	manga := Manga{}

	db.First(&manga, id)

	return manga, nil
}

func (repo MangaRepo) FindAll() ([]Manga, error) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		e := echo.New()
		e.Logger.Fatal("failed to connect database")
	}
	defer db.Close()

	mangaList := []Manga{}

	db.Find(&mangaList)

	return mangaList, nil
}

func (repo MangaRepo) Update(manga Manga) error {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		e := echo.New()
		e.Logger.Fatal("failed to connect database")
	}
	defer db.Close()

	db.Save(&manga)

	return nil
}
