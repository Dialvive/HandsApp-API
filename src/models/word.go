package models

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"id" gorm:"INT; primaryKey"`
	LocaleID       uint   `json:"locale_ID" gorm:"TINYINT NOT NULL"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"TINYINT NOT NULL"`
	Text           string `json:"name" gorm:"TEXT NOT NULL"`
	Context        string `json:"context" gorm:"TEXT"`
	Definition     string `json:"definition" gorm:"TEXT"`
	Creation       string `json:"creation" gorm:"TIMESTAMP"`
}
