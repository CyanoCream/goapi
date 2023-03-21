package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	Id     int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var bookDatas = []Book{}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": bookDatas,
	})
}
func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.Id = len(bookDatas) + 1
	bookDatas = append(bookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("bookID")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid book ID",
		})
		return
	}

	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookIndex := -1
	for i, book := range bookDatas {
		if book.Id == bookId {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found", bookId),
		})
		return
	}

	updatedBook.Id = bookId
	bookDatas[bookIndex] = updatedBook

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully updated", bookId),
	})
}

func GetBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("bookID")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid book ID",
		})
		return
	}

	bookIndex := -1
	var bookData Book

	for i, book := range bookDatas {
		if book.Id == bookId {
			bookIndex = i
			bookData = book
			break
		}
	}

	if bookIndex == -1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range bookDatas {
		if bookId == strconv.Itoa(book.Id) {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found", bookId),
		})
		return
	}

	copy(bookDatas[bookIndex:], bookDatas[bookIndex+1:])
	bookDatas[len(bookDatas)-1] = Book{}
	bookDatas = bookDatas[:len(bookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted", bookId),
	})
}
