package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalUrl string `json:"original_url" gorm:"column:original_url;not null;type:text"`
	ShortCode   string `json:"short_code" gorm:"column:short_code;uniqueIndex;not null;size:10"`
	Clicks      int    `json:"clicks" gorm:"column:clicks;default:0"`
}
