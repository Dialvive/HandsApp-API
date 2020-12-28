package models

// WordByRegion is a classification of a word by region.
type WordByRegion struct {
	WordID   uint   `json:"word_ID"`
	RegionID string `json:"region_ID" gorm:"not null"`
	Creation string `json:"creation"`
}
