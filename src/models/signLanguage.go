package models

// SignLanguage represents a Sign Language such as Mexican Sign Language.
type SignLanguage struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null"`
	Abbreviation string `json:"abbreviation" gorm:"not null"`
}
