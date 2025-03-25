package interfaces

import (
	"context"
	"simple-crud/src/model/entity"
)

type BookRepository interface {
	Create(ctx context.Context, book *entity.Book) error
	FindByID(ctx context.Context, id uint) (*entity.Book, error)
	FindAll(ctx context.Context) ([]entity.Book, error)
	Update(ctx context.Context, book *entity.Book) (*entity.Book, error)
	Delete(ctx context.Context, id uint) error
}
