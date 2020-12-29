package models

// AdCategory represents the type of ad an ad is.
type AdCategory struct {
	ID       uint   `json:"id" gorm:"TINYINT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Cost     uint   `json:"cost" gorm:"INT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}

// CreateAdCategoryInput type for ad_category POST with automatic ID.
type CreateAdCategoryInput struct {
	Name string `json:"name" binding:"required"`
	Cost uint   `json:"cost" binding:"required"`
}
