package interfaces

import (
	"simple-crud/src/model/entity"
	"simple-crud/src/model/request"
	"simple-crud/src/model/response"
)

type BookConverter interface {
	ToResponse(book *entity.Book) *response.BookResponse
	ToResponseList(books []entity.Book) []response.BookResponse
	ToEntity(req *request.BookRequest) *entity.Book
	ToUpdateEntity(req *request.BookRequestUpdate) *entity.Book
}
