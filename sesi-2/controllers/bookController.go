package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var books = []Book{}

func CreateBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.ID = fmt.Sprintf("b%d", len(books)+1)
	books = append(books, book)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Book

	err := ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, v := range books {
		if bookID == v.ID {
			condition = true
			books[i] = updatedBook
			books[i].ID = bookID

			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v is not found", bookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully updated", bookID),
	})
}

func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var book Book

	for i, v := range books {
		if bookID == v.ID {
			condition = true
			book = books[i]

			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v is not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Book": book,
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Book": books,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, v := range books {
		if bookID == v.ID {
			condition = true
			bookIndex = i

			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v is not found", bookID),
		})

		return
	}

	copy(books[bookIndex:], books[bookIndex+1:])
	books[len(books)-1] = Book{}
	books = books[:len(books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted", bookID),
	})
}
