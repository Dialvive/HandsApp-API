package models

// Region represents an administrative division of a country.
type Region struct {
	ID        uint   `json:"id" gorm:"INT; primaryKey"`
	Name      string `json:"name" gorm:"TEXT NOT NULL"`
	CountryID uint   `json:"country_ID" gorm:"TINYINT NOT NULL"`
	Creation  string `json:"creation" gorm:"TIMESTAMP"`
}
