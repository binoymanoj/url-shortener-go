package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/binoymanoj/url-shortener-go/initializers"
	"github.com/binoymanoj/url-shortener-go/models"
	"github.com/binoymanoj/url-shortener-go/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UrlController(c *gin.Context) {
	db := initializers.DB
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
	var existingURL models.URL
	result := db.Where("original_url = ?", originalUrl).First(&existingURL)

	// FIXED: Check if URL exists (no error means record was found)
	if result.Error == nil {
		// URL already exists, return existing short code
		shortURL := c.Request.Host + "/" + existingURL.ShortCode
		response := `
			<div class="result">
				<h3>✅ URL Already Exists!</h3>
				<p><strong>Original:</strong> ` + originalUrl + `</p>
				<div class="url-display">
					<input type="text" class="short-url" value="` + shortURL + `" readonly>
					<button class="copy-btn" onclick="copyToClipboard('` + shortURL + `')">📋 Copy</button>
				</div>
				<div class="stats">Clicks: ` + strconv.Itoa(existingURL.Clicks) + `</div>
			</div>
		`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
		return
	}

	// FIXED: Check if it's a "record not found" error (expected) vs other errors (unexpected)
	if result.Error != gorm.ErrRecordNotFound {
		// Unexpected database error
		log.Printf("Database error: %v", result.Error)
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(`
		<div class="error">
			<strong>Error:</strong> Database error occurred. Please try again.
		</div>
		`))
		return
	}

	// URL doesn't exist, create new short code
	shortCode := utils.GenerateShortCode(originalUrl)
	log.Printf("shortCode: %v", shortCode)

	// Ensure uniqueness
	for {
		var existing models.URL
		if db.Where("short_code = ?", shortCode).First(&existing).Error != nil {
			break
		}
		shortCode = utils.GenerateShortCode(originalUrl + strconv.FormatInt(int64(len(shortCode)), 10))
	}

	newURL := models.URL{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
	}

	if err := db.Create(&newURL).Error; err != nil {
		log.Printf("Error creating URL: %v", err)
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(`
		<div class="error">
			<strong>Error:</strong> Failed to create short URL. Please try again.
		</div>
		`))
		return
	}

	shortURL := c.Request.Host + "/" + shortCode
	response := `
	<div class="result">
		<h3>✅ URL Shortened Successfully!</h3>
		<p><strong>Original:</strong> ` + originalUrl + `</p>
		<div class="url-display">
			<input type="text" class="short-url" value="` + shortURL + `" readonly>
			<button class="copy-btn" onclick="copyToClipboard('` + shortURL + `')">📋 Copy</button>
		</div>
		<div class="stats">Clicks: 0</div>
	</div>
	`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
}
