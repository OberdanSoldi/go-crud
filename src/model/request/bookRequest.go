package request

import "time"

type BookRequest struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

type BookRequestUpdate struct {
	ID uint `json:"id"`
	BookRequest
}
