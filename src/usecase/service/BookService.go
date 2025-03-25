package service

import (
	"context"
	"errors"
	"simple-crud/src/model/request"
	"simple-crud/src/model/response"
	"simple-crud/src/usecase/interfaces"
	"strconv"
)

type bookService struct {
	repo      interfaces.BookRepository
	converter interfaces.BookConverter
}

func NewBookService(repo interfaces.BookRepository, converter interfaces.BookConverter) interfaces.BookService {
	return &bookService{
		repo:      repo,
		converter: converter,
	}
}

func (s *bookService) CreateBook(ctx context.Context, book *request.BookRequest) error {
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}

	if book.Author == "" {
		return errors.New("book author cannot be empty")
	}

	return s.repo.Create(ctx, s.converter.ToEntity(book))
}

func (s *bookService) GetBookByID(ctx context.Context, id string) (*response.BookResponse, error) {
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	book, err := s.repo.FindByID(ctx, uint(parsedId))
	if err != nil {
		return nil, err
	}

	return s.converter.ToResponse(book), nil
}

func (s *bookService) GetAllBooks(ctx context.Context) ([]response.BookResponse, error) {
	books, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return s.converter.ToResponseList(books), nil
}

func (s *bookService) UpdateBook(ctx context.Context, book *request.BookRequestUpdate, id string) (*response.BookResponse, error) {
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return nil, errors.New("invalid ID format")
	}

	existingBook, err := s.repo.FindByID(ctx, book.ID)
	if err != nil {
		return nil, err
	}

	if existingBook == nil {
		return nil, errors.New("book not found")
	}

	if book.Title == "" {
		return nil, errors.New("book title cannot be empty")
	}

	if book.Author == "" {
		return nil, errors.New("book author cannot be empty")
	}

	convertedBook := s.converter.ToUpdateEntity(book)

	updatedBook, err := s.repo.Update(ctx, convertedBook)
	if err != nil {
		return nil, err
	}

	return s.converter.ToResponse(updatedBook), nil
}

func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("invalid ID format")
	}

	existingBook, _ := s.repo.FindByID(ctx, uint(parsedId))
	if existingBook == nil {
		return errors.New("book not found")
	}

	return s.repo.Delete(ctx, uint(parsedId))
}
