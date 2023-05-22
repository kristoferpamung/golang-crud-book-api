package main

import (
	"book_api/controllers"
	"book_api/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/books", controllers.GetAllBooks)
	v1.POST("/books", controllers.CreateBook)
	v1.GET("/books/:id", controllers.GetBookById)
	v1.PATCH("/books/:id", controllers.UpdateBook)
	v1.DELETE("/books/:id", controllers.DeleteBook)

	router.Run(os.Getenv("SERVER_PORT"))
}
