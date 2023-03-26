package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-web-services/model"
)

var books = []model.Book{}

// @BasePath /api/v1

// CreateBook godoc
// @Summary Create a single book
// @Schemes
// @Description Create a single book by providing title, author, and desc
// @Accept json
// @Produce json
// @Param data body model.Book true "Sample payload"
// @Success 200 {object} response.JSONSuccessResult{data=model.Book,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /book [post]
func CreateBook(ctx *gin.Context) {
	var book model.Book

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

// @BasePath /api/v1

// UpdateBook godoc
// @Summary Update a single book
// @Schemes
// @Description Update a single book by providing title, author, and desc
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param data body model.Book true "Sample payload"
// @Success 200 {object} response.JSONSuccessResult{data=model.Book,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /book/{id} [put]
func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook model.Book

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

// @BasePath /api/v1

// GetBookByID godoc
// @Summary Get a single book
// @Schemes
// @Description Get a single book by providing its ID
// @Accept json
// @Produce json
// @Param id path string true "Sample payload"
// @Success 200 {object} response.JSONSuccessResult{data=model.Book,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /book/{id} [get]
func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var book model.Book

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

// @BasePath /api/v1

// GetBooks godoc
// @Summary Get multiple books
// @Schemes
// @Description Create multiple books
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONSuccessResult{data=model.Book,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /books [get]
func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Book": books,
	})
}

// @BasePath /api/v1

// DeleteBook godoc
// @Summary Delete a single book
// @Schemes
// @Description Delete a single book by providing its ID
// @Accept json
// @Produce json
// @Param id path string true "Sample payload"
// @Success 200 {object} response.JSONSuccessResult{data=model.Book,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /book/{id} [delete]
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
	books[len(books)-1] = model.Book{}
	books = books[:len(books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted", bookID),
	})
}
