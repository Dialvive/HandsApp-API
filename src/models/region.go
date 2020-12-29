package models

// Region represents an administrative division of a country.
type Region struct {
	ID        uint   `json:"id" gorm:"INT; primaryKey"`
	Name      string `json:"name" gorm:"TEXT NOT NULL"`
	CountryID uint   `json:"country_ID" gorm:"TINYINT NOT NULL"`
	Creation  string `json:"creation" gorm:"TIMESTAMP"`
}

// CreateRegionInput represents an administrative division of a country.
type CreateRegionInput struct {
	Name      string `json:"name" binding:"required"`
	CountryID int    `json:"country_ID" binding:"required"`
}
