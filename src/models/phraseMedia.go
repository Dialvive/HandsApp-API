package models

// PhraseMedia stores the URL of the phrase's media.
type PhraseMedia struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	URL         string `json:"url"`
	Description string `json:"description" gorm:"not null"`
	Creation    string `json:"creation"`
}
