package controllers

import (
	"net/http"

	"github.com/binoymanoj/url-shortener-go/initializers"
	"github.com/binoymanoj/url-shortener-go/models"
	"github.com/gin-gonic/gin"
)

func RedirectionController(c *gin.Context) {
	shortCode := c.Param("shortcode")
	db := initializers.DB

	var url models.URL

	if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"message": "Short URL not found",
		})

		return
	}

	db.Model(&url).Update("clicks", url.Clicks+1)

	c.Redirect(http.StatusFound, url.OriginalUrl)
}
