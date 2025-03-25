package interfaces

import (
	"context"
	"simple-crud/src/model/request"
	"simple-crud/src/model/response"
)

type BookService interface {
	CreateBook(ctx context.Context, book *request.BookRequest) error
	GetBookByID(ctx context.Context, id string) (*response.BookResponse, error)
	GetAllBooks(ctx context.Context) ([]response.BookResponse, error)
	UpdateBook(ctx context.Context, book *request.BookRequestUpdate, id string) (*response.BookResponse, error)
	DeleteBook(ctx context.Context, id string) error
}
