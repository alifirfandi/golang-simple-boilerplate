package model

import "time"

type BookResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int    `json:"year"`
}
