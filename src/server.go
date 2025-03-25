package src

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-crud/src/adapter/controller"
	"simple-crud/src/adapter/db"
	"simple-crud/src/adapter/repository"
	"simple-crud/src/usecase/converter"
	"simple-crud/src/usecase/service"
)

func StartServer() {
	db.InitDb()
	dbConn := db.GetDb()

	bookRepo := repository.NewBookRepository(dbConn)
	bookConverter := converter.NewBookConverter()
	bookService := service.NewBookService(bookRepo, bookConverter)
	bookController := controller.NewBookController(bookService)

	router := gin.Default()

	bookController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
