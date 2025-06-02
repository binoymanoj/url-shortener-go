package main

import (
	"net/http"
	"os"

	"github.com/binoymanoj/url-shortener-go/controllers"
	"github.com/binoymanoj/url-shortener-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", controllers.HomeController)
	r.POST("/shrtnurl", controllers.UrlController)
	r.GET("/:shortcode", controllers.RedirectionController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	r.Run(":" + port)
}
