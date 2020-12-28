package models

// Phrase represents a coherent spoken language set of words.
type Phrase struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	LocaleID         uint   `json:"locale_ID" gorm:"not null"`
	MediaID          uint   `json:"media_ID" gorm:"not null"`
	PhraseCategoryID uint   `json:"phrase_category_ID" gorm:"not null"`
	Text             string `json:"name" gorm:"not null"`
	Context          string `json:"context"`
	PhraseCategory   uint   `json:"phraseCategoryID" gorm:"not null"`
	Creation         string `json:"creation"`
}
