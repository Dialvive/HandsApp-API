package models

// SpokenLanguage represents a language such as English, Spanish, etc.
type SpokenLanguage struct {
	ID           uint   `json:"id" gorm:"INT; primaryKey"`
	Name         string `json:"name" gorm:"TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"VARCHAR(2) NOT NULL"`
	Creation     string `json:"creation" gorm:"TIMESTAMP"`
}

// CreateSpokenLanguageInput type with automatic ID.
type CreateSpokenLanguageInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}
