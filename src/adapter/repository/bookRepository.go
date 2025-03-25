package repository

import (
	"context"
	"gorm.io/gorm"
	"simple-crud/src/model/entity"
	"simple-crud/src/usecase/interfaces"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(ctx context.Context, book *entity.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *bookRepository) FindByID(ctx context.Context, id uint) (*entity.Book, error) {
	var book entity.Book
	err := r.db.WithContext(ctx).First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepository) FindAll(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.WithContext(ctx).Find(&books).Error
	return books, err
}

func (r *bookRepository) Update(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	err := r.db.WithContext(ctx).Save(book)
	if err.Error != nil {
		return nil, err.Error
	}

	return book, nil
}

func (r *bookRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Book{}, id).Error
}
