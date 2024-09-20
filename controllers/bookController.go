package controllers

import (
	"net/http"
	"strconv"

	"github.com/f1nn-ach/go-jwt/initializiers"
	"github.com/f1nn-ach/go-jwt/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializiers.LoadEnvVariables()
	initializiers.ConnectToDb()
	initializiers.SyncDb()
}

func CreateBook(c *gin.Context) {
	bookName := c.PostForm("bookName")
	urlImage := c.PostForm("urlImage")
	priceStr := c.PostForm("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price value"})
		return
	}

	book := models.Book{
		BookName: bookName,
		UrlImage: urlImage,
		Price:    price,
	}

	result := initializiers.DB.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully!",
	})

}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	result := initializiers.DB.Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No have book in database"})
		return
	}

	c.HTML(http.StatusOK, "book_view.html", gin.H{
		"Books": books,
	})
}

func DeleteBook(c *gin.Context) {
	var id = c.Param("ID")
	var book models.Book

	result := initializiers.DB.Where("ID = ?", id).Delete(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No book found with that ID"})
		return
	}

	var books []models.Book
	initializiers.DB.Find(&books)

	c.HTML(http.StatusOK, "book_view.html", gin.H{
		"Books": books,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("ID")
	var book models.Book

	bookName := c.PostForm("bookName")
	urlImage := c.PostForm("urlImage")
	priceStr := c.PostForm("price")

	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price value"})
		return
	}

	result := initializiers.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	book.BookName = bookName
	book.UrlImage = urlImage
	book.Price = priceFloat

	if err := initializiers.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	var books []models.Book
	initializiers.DB.Find(&books)

	c.HTML(http.StatusOK, "book_view.html", gin.H{
		"Books": books,
	})
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "create_page.html", nil)
}

func GetEdit(c *gin.Context) {
	var id = c.Param("ID")
	var book models.Book

	result := initializiers.DB.Where("ID = ?", id).Find(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No book found with that ID"})
		return
	}

	c.HTML(http.StatusOK, "edit_book.html", book)
}

// func CreateBook(c *gin.Context) {
// 	book := models.Book{
// 		BookName: "ตำนานดาบและคทาแห่งวิสตอเรีย 1",
// 		UrlImage: "https://cdn-local.mebmarket.com/meb/server1/189057/Thumbnail/book_detail_large.gif?3",
// 		Price:    100,
// 	}
// 	result := initializiers.DB.Create(&book)

// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"massage": book,
// 	})
// }

// func GetBook(c *gin.Context) {
// 	bookID := c.Param("id")

// 	var book models.Book
// 	result := initializiers.DB.First(&book, bookID)

// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.HTML(http.StatusOK, "book_view.html", gin.H{
// 		"BookName": book.BookName,
// 		"UrlImage": book.UrlImage,
// 		"Price":    book.Price,
// 	})
// }
