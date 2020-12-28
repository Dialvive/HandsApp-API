package models

// PhraseCategory is a category in which a set of phrases fall into.
type PhraseCategory struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Creation string `json:"creation"`
}
