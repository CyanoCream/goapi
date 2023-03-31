package service

import "sesi_8/model"

type BookService interface {
	CreateBook(model.Books) (model.Books, error)
	GetBooks() ([]model.Books, error)
	GetBook(id int) (model.Books, error)
	UpdateBook(book model.Books) (model.Books, error)
	DeleteBook(id int) error
}

func (s *Service) CreateBook(book model.Books) (model.Books, error) {
	books, err := s.repo.CreateBook(book)
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *Service) GetBooks() ([]model.Books, error) {
	book, err := s.repo.GetBooks()
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *Service) GetBook(id int) (model.Books, error) {
	book, err := s.repo.GetBook(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *Service) UpdateBook(book model.Books) (model.Books, error) {
	book, err := s.repo.UpdateBook(book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *Service) DeleteBook(id int) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
