package model

import "time"

// Manga is define of manga object for gorm
type Manga struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	ImgURL    string    `json:"img"`
	Number    int       `json:"readNumber"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
