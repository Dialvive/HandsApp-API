package models

// WordMedia stores the URL of the word's media.
type WordMedia struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	URL         string `json:"url"`
	Description string `json:"description" gorm:"not null"`
	Creation    string `json:"creation"`
}
