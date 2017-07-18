package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MangaRepoInterface define what MangaRepo struct have some methods.
type MangaRepoInterface interface {
	Insert(title string, imgURL string) error
	Find(id int) (Manga, error)
	FindAll() ([]Manga, error)
	Update(Manga) error
}

type MangaRepo struct{}

func NewMangaRepo() MangaRepoInterface {
	return MangaRepo{}
}

func (repo MangaRepo) Insert(title string, imgURL string) error {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("failed to connect database")
		return err
	}
	defer db.Close()

	db.Create(&Manga{Title: title, ImgURL: imgURL, Number: 1})

	return nil
}

// Find is Method that get Manga by id
func (repo MangaRepo) Find(id int) (Manga, error) {
	manga := Manga{}

	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("failed to connect database")
		return manga, err
	}
	defer db.Close()

	db.First(&manga, id)

	return manga, nil
}

// FindAll is Method that get all Manga
func (repo MangaRepo) FindAll() ([]Manga, error) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("failed to connect database")
		return nil, err
	}
	defer db.Close()

	mangaList := []Manga{}

	db.Find(&mangaList)

	return mangaList, nil
}

func (repo MangaRepo) Update(manga Manga) error {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/manga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("failed to connect database")
		return err
	}
	defer db.Close()

	db.Save(&manga)

	return nil
}
