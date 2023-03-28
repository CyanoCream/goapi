package controllers

import (
	"Latihan1/config"
	"Latihan1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllBooks(c *gin.Context) {
	books := []models.Book{}
	rows, err := config.Db.Query("SELECT id, title, author, description FROM book")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting books",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error scanning books",
			})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
		})
		return
	}

	book := models.Book{}
	err = config.Db.QueryRow("SELECT id, title, author, description FROM book WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting book",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO book (title, author, description) VALUES ($1, $2, $3) RETURNING id`
	err := config.Db.QueryRow(sqlStatement, book.Title, book.Author, book.Description).Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": book})
}
func UpdateBook(c *gin.Context) {
	var book models.Book

	// bind request body ke struct book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ambil id buku dari URL param
	id := c.Param("id")

	// update data buku di database
	sqlStatement := `UPDATE book SET title=$2, author=$3, description=$4 WHERE id=$1`
	res, err := config.Db.Exec(sqlStatement, id, book.Title, book.Author, book.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// cek apakah data berhasil diupdate
	count, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	// Check if the book exists
	var book models.Book
	err := config.Db.QueryRow("SELECT * FROM book WHERE id = $1", bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting book",
		})
		return
	}

	// Delete the book
	_, err = config.Db.Exec("DELETE FROM book WHERE id = $1", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted"),
	})
}
