package models

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"id" gorm:"INT; primaryKey"`
	LocaleID       uint   `json:"locale_ID" gorm:"TINYINT NOT NULL"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"TINYINT NOT NULL"`
	Text           string `json:"name" gorm:"TEXT NOT NULL"`
	Context        string `json:"context" gorm:"TEXT"`
	Definition     string `json:"definition" gorm:"TEXT"`
	Modified       string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateWordInput represents a coherent spoken word.
type CreateWordInput struct {
	LocaleID       uint   `json:"locale_ID" binding:"required"`
	WordCategoryID uint   `json:"phrase_category_ID" binding:"required"`
	Text           string `json:"name" binding:"required"`
	Definition     string `json:"definition" binding:"required"`
	Context        string `json:"context" binding:"required"`
}
