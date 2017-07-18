package request

// MangaCreate is post request parameter of Manga Item Creating
type MangaCreate struct {
	Title  string `json:"title" validate:"required"`
	ImgURL string `json:"img_url" validate:"required"`
	Number int    `json:"number"`
}

// MangaUpdate is post request parameter of Manga Item Updating
type MangaUpdate struct {
	Title  string `json:"title"`
	ImgURL string `json:"img_url"`
	Number int    `json:"number"`
}
