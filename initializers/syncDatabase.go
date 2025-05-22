package initializers

import "github.com/binoymanoj/url-shortener-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Url{})
}
