package models

// Locale is a mix of location, spoken language, and sign language classification..
type Locale struct {
	ID               uint   `json:"id" gorm:"SMALLINT; primaryKey"`
	CountryID        uint   `json:"country_ID" gorm:"TINYINT NOT NULL"`
	SpokenLanguageID uint   `json:"spoken_language_ID" gorm:"TINYINT NOT NULL"`
	SignLanguageID   uint   `json:"sign_language_ID" gorm:"TINYINT NOT NULL"`
	Creation         string `json:"creation" gorm:"TIMESTAMP"`
}

// CreateLocaleInput is a mix of location, spoken language, and sign language classification..
type CreateLocaleInput struct {
	CountryID        uint   `json:"country_ID" binding:"required"`
	SpokenLanguageID uint   `json:"spoken_language_ID" binding:"required"`
	SignLanguageID   uint   `json:"sign_language_ID" binding:"required"`
	Creation         string `json:"creation" binding:"required"`
}
