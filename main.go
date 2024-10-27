package main

import (
	"book-crud/db"
	"book-crud/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Book Processing Functions

func GetBooks(c *gin.Context) {
	var books []models.Book
	if result := db.DB.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if result := db.DB.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if result := db.DB.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if result := db.DB.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if result := db.DB.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "nu privet. Nachinay"})
}

func main() {

	db.Init() //Initializing the database

	router := gin.Default()

	router.GET("/", RootHandler)
	router.GET("/books", GetBooks)
	router.POST("/books", CreateBook)
	router.GET("/books/:id", GetBookByID)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)

	log.Fatal(router.Run(":8081"))
}
