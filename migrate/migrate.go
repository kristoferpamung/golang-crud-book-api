package main

import (
	"book_api/initializers"
	"book_api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
