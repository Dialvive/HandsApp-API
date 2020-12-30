package models

// WordByRegion is a classification of a word by region.
type WordByRegion struct {
	WordID   uint   `json:"word_ID" gorm:"INT NOT NULL"`
	RegionID string `json:"region_ID" gorm:"INT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateWordByRegionInput is a classification of a word by region.
type CreateWordByRegionInput struct {
	WordID   uint   `json:"word_ID" binding:"required"`
	RegionID string `json:"region_ID" binding:"required"`
}
