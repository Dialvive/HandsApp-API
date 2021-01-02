package models

// Advertisement represents an ad.
type Advertisement struct {
	ID           uint   `json:"id" gorm:"INT AUTO_INCREMENT; primaryKey"`
	UserID       uint   `json:"user_ID" gorm:"INT NOT NULL"`
	RegionID     uint   `json:"region_ID" gorm:"INT NOT NULL"`
	AdCategoryID uint   `json:"ad_category_ID" gorm:"TINYINT NOT NULL"`
	Title        string `json:"title" gorm:"VARCHAR(64) NOT NULL"`
	Body         string `json:"body" gorm:"TEXT"`
	Media        bool   `json:"media" gorm:"BOOLEAN; default:0"`
	Paid         uint   `json:"paid" gorm:"INT NOT NULL"`
	Modified     string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateAdvertisementInput represents an ad.
type CreateAdvertisementInput struct {
	UserID       uint   `json:"user_ID" binding:"required"`
	RegionID     uint   `json:"region_ID" binding:"required"`
	AdCategoryID uint   `json:"ad_category_ID" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Body         string `json:"body" binding:"required"`
	Media        bool   `json:"media" binding:"required"`
	Paid         uint   `json:"paid" binding:"required"`
}
