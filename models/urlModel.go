package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex;not null"`
	Clicks      int    `gorm:"default:0"`
}
