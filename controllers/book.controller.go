package controllers

import (
	"book_api/initializers"
	"book_api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Author      string  `json:"Author" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
	Price       uint    `json:"price" binding:"required"`
	ImageUrl    string  `json:"image_url"`
}

func GetAllBooks(ctx *gin.Context) {
	var books []models.Book

	initializers.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   books,
	})
}

func CreateBook(ctx *gin.Context) {
	var input CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		Author:      input.Author,
		Rating:      input.Rating,
		Price:       input.Price,
		ImageUrl:    input.ImageUrl,
	}

	initializers.DB.Create(&book)

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   book,
	})

}

func GetBookById(ctx *gin.Context) {
	var book models.Book

	id := ctx.Param("id")

	if err := initializers.DB.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   book,
	})
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book
	var input CreateBookInput

	id := ctx.Param("id")

	if err := initializers.DB.Where("id =?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found!",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	initializers.DB.Model(&book).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   book,
	})
}

func DeleteBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")

	if err := initializers.DB.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found!",
		})
	}

	initializers.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success delete book with id : %s", id),
	})

}
