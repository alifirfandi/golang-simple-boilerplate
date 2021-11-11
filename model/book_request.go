package model

type BookRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
}
