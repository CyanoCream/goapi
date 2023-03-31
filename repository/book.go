package repository

import (
	"database/sql"
	"fmt"
	"sesi_8/model"
)

// interface employee
type BookRepo interface {
	CreateBook(book model.Books) (model.Books, error)
	GetBooks() ([]model.Books, error)
	GetBook(id int) (model.Books, error)
	UpdateBook(book model.Books) (model.Books, error)
	DeleteBook(id int) error
}

func (r Repo) CreateBook(book model.Books) (model.Books, error) {
	sqlStatement := `
	INSERT INTO book(title, author, description)
	VALUES($1, $2, $3)
	RETURNING *
	`
	newBook := model.Books{}

	stm, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return newBook, err
	}

	err = stm.QueryRow(book.Title, book.Author, book.Description).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (r Repo) GetBooks() ([]model.Books, error) {
	sqlStatement := `
	SELECT id, title, author, description FROM book 
	`
	books := []model.Books{}

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Books{}

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r Repo) GetBook(id int) (model.Books, error) {
	sqlStatement := `
	SELECT id, title, author, description FROM book
	WHERE id = $1
	`
	book := model.Books{}

	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err == sql.ErrNoRows {
		return book, fmt.Errorf("book with ID %d not found", id)
	} else if err != nil {
		return book, err
	} else {
		return book, nil
	}
}

func (r Repo) UpdateBook(book model.Books) (model.Books, error) {
	sqlStatement := `
	UPDATE book SET title = $2, author = $3, description = $4
	WHERE id = $1
	RETURNING *
	`
	newBook := model.Books{}

	stm, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return newBook, err
	}

	err = stm.QueryRow(book.ID, book.Title, book.Author, book.Description).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (r Repo) DeleteBook(id int) error {
	sqlStatement := `
		DELETE FROM book WHERE id = $1
	`

	res, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count != 0 {
		return nil
	}

	return fmt.Errorf("book with ID %d not found", id)
}
