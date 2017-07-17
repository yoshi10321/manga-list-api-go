package model

import "github.com/jinzhu/gorm"

// Manga is define of manga object
type Manga struct {
	gorm.Model
	Title  string `json:"title"`
	ImgURL string `json:"img_url"`
	Number int    `json:"number"`
}
