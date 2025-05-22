package main

import (
	"github.com/binoymanoj/url-shortener-go/initializers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "you reached ~",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/shrtnurl")

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	r.Run(":" + port)
}
