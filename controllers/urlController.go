package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UrlController(c *gin.Context) {
	originalUrl := c.PostForm("url")

	if originalUrl == "" {
		c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(`
		<div class="error">
			<strong>Error:</strong> Please provide a valid URL
		</div>
		`))

		return
	}

	// checking for protocol
	if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
		originalUrl = "https://" + originalUrl
	}

	// checking if url already exists

}
