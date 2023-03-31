package handler

import (
	"net/http"
	"sesi_8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	bookRequest := model.Books{}
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = h.app.CreateBook(bookRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Created",
	})
}

func (h HttpServer) GetBooks(c *gin.Context) {
	books, err := h.app.GetBooks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    books,
	})
}

func (h HttpServer) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := h.app.GetBook(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    book,
	})
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookRequest := model.Books{}
	err = c.ShouldBindJSON(&bookRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookRequest.ID = id

	_, err = h.app.CreateBook(bookRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Updated",
	})
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.app.DeleteBook(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}
