package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-crud/src/model/request"
	"simple-crud/src/usecase/interfaces"
)

type bookController struct {
	bookService interfaces.BookService
}

func NewBookController(bookService interfaces.BookService) interfaces.BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c *bookController) RegisterRoutes(router *gin.Engine) {
	bookGroup := router.Group("/api/books")
	{
		bookGroup.POST("", c.CreateBook)
		bookGroup.GET("", c.GetAllBooks)
		bookGroup.GET("/:id", c.GetBookByID)
		bookGroup.PUT("/:id", c.UpdateBook)
		bookGroup.DELETE("/:id", c.DeleteBook)
	}
}

func (c *bookController) CreateBook(ctx *gin.Context) {
	var bookReq *request.BookRequest
	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.bookService.CreateBook(ctx, bookReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
	})
}

func (c *bookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.bookService.GetAllBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (c *bookController) GetBookByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	book, err := c.bookService.GetBookByID(ctx, idParam)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var bookReq *request.BookRequestUpdate
	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := c.bookService.UpdateBook(ctx, bookReq, idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"data":    updatedBook,
	})
}

func (c *bookController) DeleteBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.bookService.DeleteBook(ctx, idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
