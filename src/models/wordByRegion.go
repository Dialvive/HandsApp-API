package models

// WordByRegion is a classification of a word by region.
type WordByRegion struct {
	WordID   uint   `json:"word_ID" gorm:"INT NOT NULL"`
	RegionID string `json:"region_ID" gorm:"INT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}
