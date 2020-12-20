package models

// Word represents an atomic unit of a spoken language.
type Word struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Text         string `json:"name" gorm:"not null"`
	Media        string `json:"media"`
	Definition   string `json:"definition"`
	Context      string `json:"context"`
	WordCategory uint   `json:"wordCategoryID" gorm:"not null"`
}
