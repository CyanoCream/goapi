package repository

import (
	"fmt"
	"sesi_8/model"
)

type BookRepo interface {
	CreateBook(book model.Books) (model.Books, error)
	GetBooks() ([]model.Books, error)
	GetBook(id int) (model.Books, error)
	UpdateBook(book model.Books) (model.Books, error)
	DeleteBook(id int) error
}

func (r Repo) CreateBook(book model.Books) (model.Books, error) {
	newBook := model.Books{}
	result := r.db.Create(&book)
	if result.Error != nil {
		return newBook, result.Error
	}
	return book, nil
}

func (r Repo) GetBooks() ([]model.Books, error) {
	var books []model.Books
	result := r.db.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (r Repo) GetBook(id int) (model.Books, error) {
	var book model.Books
	result := r.db.First(&book, id)
	if result.Error != nil {
		return book, fmt.Errorf("book with ID %d not found", id)
	}
	return book, nil
}

func (r Repo) UpdateBook(book model.Books) (model.Books, error) {
	result := r.db.Updates(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (r Repo) DeleteBook(id int) error {
	result := r.db.Delete(&model.Books{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("book with ID %d not found", id)
	}
	return nil
}
