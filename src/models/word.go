package models

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	LocaleID       uint   `json:"locale_ID" gorm:"not null"`
	MediaID        uint   `json:"media_ID" gorm:"not null"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"not null"`
	Text           string `json:"name" gorm:"not null"`
	Context        string `json:"context"`
	Definition     string `json:"definition"`
	Creation       string `json:"creation"`
}
