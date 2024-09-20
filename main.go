package main

import (
	"github.com/f1nn-ach/go-jwt/controllers"
	"github.com/f1nn-ach/go-jwt/initializiers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializiers.LoadEnvVariables()
	initializiers.ConnectToDb()
	initializiers.SyncDb()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/html/*")

	r.GET("/", controllers.GetIndex)

	r.GET("/get", controllers.GetAllBooks)
	r.GET("/create", controllers.GetCreate)

	r.POST("/api/create-book", controllers.CreateBook)
	r.GET("/delete/:ID", controllers.DeleteBook)
	r.GET("/edit/:ID", controllers.GetEdit)

	r.POST("/api/update-book", controllers.UpdateBook)

	r.Run()
}
