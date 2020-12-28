package models

// Region represents an administrative division of a country.
type Region struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"not null"`
	CountryID uint   `json:"country_ID" gorm:"not null"`
	Creation  string `json:"creation"`
}
