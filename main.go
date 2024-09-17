package main

import (
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pingpong",
		})
	})
	r.Run()
}
