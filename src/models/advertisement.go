package models

// Advertisement represents an ad.
type Advertisement struct {
	ID           uint   `json:"id" gorm:"INT AUTO_INCREMENT; primaryKey"`
	UserID       uint   `json:"user_ID" gorm:"INT NOT NULL"`
	RegionID     uint   `json:"region_ID" gorm:"INT NOT NULL"`
	AdCategoryID uint   `json:"ad_category_ID" gorm:"TINYINT NOT NULL"`
	Title        string `json:"title" gorm:"VARCHAR(64) NOT NULL"`
	Body         string `json:"body" gorm:"TEXT"`
	Media        bool   `json:"media" gorm:"BOOLEAN; default: false"`
	Paid         uint   `json:"paid" gorm:"INT NOT NULL"`
	Creation     string `json:"creation" gorm:"TIMESTAMP"`
}
