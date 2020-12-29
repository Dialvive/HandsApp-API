package models

// Country represents a real world country.
type Country struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null"`
	Abbreviation string `json:"abbreviation" gorm:"not null"`
	Creation     string `json:"creation"`
}

// CreateCountryInput type for country POST with automatic ID.
type CreateCountryInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// UpdateCountryInput type for country POST without bindings.
type UpdateCountryInput struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}
