package models

// WordCategory is a category in which a set of words fall into.
type WordCategory struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}
