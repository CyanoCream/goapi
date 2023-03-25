package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "perpus"
)

var db *sql.DB

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully!")

	router := gin.Default()

	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", deleteBook)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getAllBooks(c *gin.Context) {
	books := []Book{}
	rows, err := db.Query("SELECT id, title, author, description FROM book")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting books",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := Book{}
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

func getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
		})
		return
	}

	book := Book{}
	err = db.QueryRow("SELECT id, title, author, description FROM book WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting book",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO book (title, author, description) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, book.Title, book.Author, book.Description).Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": book})
}
func UpdateBook(c *gin.Context) {
	var book Book

	// bind request body ke struct book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ambil id buku dari URL param
	id := c.Param("id")

	// update data buku di database
	sqlStatement := `UPDATE book SET title=$2, author=$3, description=$4 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, book.Title, book.Author, book.Description)
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
func deleteBook(c *gin.Context) {
	bookID := c.Param("id")

	// Check if the book exists
	var book Book
	err := db.QueryRow("SELECT * FROM book WHERE id = $1", bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting book",
		})
		return
	}

	// Delete the book
	_, err = db.Exec("DELETE FROM book WHERE id = $1", bookID)
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
