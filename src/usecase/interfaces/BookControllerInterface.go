package interfaces

import "github.com/gin-gonic/gin"

type BookController interface {
	RegisterRoutes(router *gin.Engine)
	CreateBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
	GetBookByID(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}
