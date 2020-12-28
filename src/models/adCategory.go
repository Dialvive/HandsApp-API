package models

// AdCategory represents the type of ad an ad is.
type AdCategory struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Cost     string `json:"cost" gorm:"not null"`
	Creation string `json:"creation"`
}
