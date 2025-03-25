package converter

import (
	"simple-crud/src/model/entity"
	"simple-crud/src/model/request"
	"simple-crud/src/model/response"
	"simple-crud/src/usecase/interfaces"
)

type bookConverter struct{}

func NewBookConverter() interfaces.BookConverter {
	return &bookConverter{}
}

func (c *bookConverter) ToResponse(book *entity.Book) *response.BookResponse {
	return &response.BookResponse{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: book.PublishedDate,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}
}

func (c *bookConverter) ToEntity(req *request.BookRequest) *entity.Book {
	book := entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	return &book
}

func (c *bookConverter) ToResponseList(books []entity.Book) []response.BookResponse {
	responses := make([]response.BookResponse, len(books))

	for i, book := range books {
		bookCopy := book
		bookResponse := c.ToResponse(&bookCopy)
		responses[i] = *bookResponse
	}

	return responses
}

func (c *bookConverter) ToUpdateEntity(req *request.BookRequestUpdate) *entity.Book {
	book := entity.Book{
		ID:            req.ID,
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: req.PublishedDate,
	}

	return &book
}
