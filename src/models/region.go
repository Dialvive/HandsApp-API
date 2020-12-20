package models

// Region represents an administrative division of a country.
type Region struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"not null"`
	Country uint   `json:"country" gorm:"not null"`
}
