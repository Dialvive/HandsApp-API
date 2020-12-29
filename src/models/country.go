package models

// Country represents a real world country.
type Country struct {
	ID           uint   `json:"id" gorm:"TINYINT; primaryKey"`
	Name         string `json:"name" gorm:"TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"VARCHAR(2) NOT NULL"`
	Creation     string `json:"creation" gorm:"TIMESTAMP"`
}

// CreateCountryInput type for country POST with automatic ID.
type CreateCountryInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}
