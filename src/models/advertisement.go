package models

// Advertisement represents an ad.
type Advertisement struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	UserID       uint   `json:"user_ID" gorm:"not null"`
	RegionID     uint   `json:"region_ID" gorm:"not null"`
	AdCategoryID uint   `json:"ad_category_ID" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Body         string `json:"body"`
	Media        bool   `json:"media"`
	Paid         uint   `json:"paid" gorm:"not null"`
	Creation     string `json:"creation"`
}
