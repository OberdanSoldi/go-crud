package response

import "time"

type BookResponse struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PublishedDate time.Time `json:"published_date"`
}
