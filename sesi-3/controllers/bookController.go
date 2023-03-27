package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-web-services/config"
	"github.com/nurhidaylma/dts-web-services/model"
)

var books = []model.Book{}
var db *sql.DB

func init() {
	db = config.DB()
}

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
	var bookLen int

	sqlSelectStatement := `SELECT COUNT(*) FROM books`
	sqlInsertStatement := `
		INSERT INTO books (id, title, author, description)
		VALUES ($1, $2, $3, $4)
		Returning *
	`

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	// Get books length
	err := db.QueryRow(sqlSelectStatement).Scan(&bookLen)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Database Error",
			"error_message": fmt.Sprintf("%s", err),
		})

		return
	}
	fmt.Println(*&bookLen)
	book.ID = fmt.Sprintf("b%d", bookLen+1)

	// Insert into books
	err = db.QueryRow(sqlInsertStatement, book.ID, book.Title, book.Author, book.Desc).
		Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Database Error",
			"error_message": fmt.Sprintf("%s", err),
		})

		return
	}

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
	var updatedBook model.Book
	bookID := ctx.Param("bookID")

	sqlStatement := `
		UPDATE books 
		SET title = $2, author = $3, description = $4
		WHERE id = $1
	`
	err := ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = db.Exec(sqlStatement, bookID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Database Error",
			"error_message": fmt.Sprintf("%s", err),
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
	var book model.Book
	bookID := ctx.Param("bookID")
	condition := true

	sqlStatement := `SELECT * FROM books WHERE id = $1`
	err := db.QueryRow(sqlStatement, bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		condition = false
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
	sqlStatement := `SELECT * FROM books`
	row, err := db.Query(sqlStatement)
	for row.Next() {
		var book model.Book
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error_status":  "Error Database",
				"error_message": fmt.Sprintf("%s", err),
			})

			return
		}

		books = append(books, book)
	}

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
	var condition bool

	sqlStatement := `DELETE from books WHERE id = $1`

	_, err := db.Exec(sqlStatement, bookID)
	if err != nil {
		condition = true
	}

	if condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v is not found", bookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted", bookID),
	})
}
