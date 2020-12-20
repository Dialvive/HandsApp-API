package models

// Phrase represents a coherent spoken language set of words.
type Phrase struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Text           string `json:"name" gorm:"not null"`
	Media          string `json:"media"`
	Context        string `json:"context"`
	PhraseCategory uint   `json:"phraseCategoryID" gorm:"not null"`
}
