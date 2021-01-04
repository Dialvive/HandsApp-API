package models

import "github.com/dgrijalva/jwt-go"

// Advertisement represents an ad.
type Advertisement struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL; primaryKey"`
	UserID       uint   `json:"user_ID" gorm:"column:user_ID;type:INT NOT NULL"`
	RegionID     uint   `json:"region_ID" gorm:"column:region_ID;type:INT NOT NULL"`
	AdCategoryID uint   `json:"ad_category_ID" gorm:"column:ad_category_ID;type:TINYINT NOT NULL"`
	Title        string `json:"title" gorm:"column:title;type:VARCHAR(64) NOT NULL"`
	Body         string `json:"body" gorm:"column:body;type:TEXT"`
	Media        bool   `json:"media" gorm:"column:media;BOOLEAN; default:0"`
	Paid         uint   `json:"paid" gorm:"type:column:paid;INT NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
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

// AdvertisementClaim is a claim that cointains Advertisement as Data.
type AdvertisementClaim struct {
	Data Advertisement `json:"data" binding:"required"`
	jwt.StandardClaims
}
