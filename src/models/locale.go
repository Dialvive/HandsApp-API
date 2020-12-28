package models

// Locale is a mix of location, spoken language, and sign language classification..
type Locale struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	CountryID        uint   `json:"country_ID" gorm:"not null"`
	SpokenLanguageID uint   `json:"spoken_language_ID" gorm:"not null"`
	SignLanguageID   uint   `json:"sign_language_ID" gorm:"not null"`
	Creation         string `json:"creation"`
}
