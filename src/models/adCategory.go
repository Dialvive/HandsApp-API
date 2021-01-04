package models

import "github.com/dgrijalva/jwt-go"

// AdCategory represents the type of ad an ad is.
type AdCategory struct {
	ID       uint   `json:"ID" gorm:"type:TINYINT NOT NULL;primaryKey"`
	Name     string `json:"name" gorm:"type:TEXT NOT NULL"`
	Cost     uint   `json:"cost" gorm:"type:INT NOT NULL"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateAdCategoryInput type for ad_category POST with automatic ID.
type CreateAdCategoryInput struct {
	Name string `json:"name" binding:"required"`
	Cost uint   `json:"cost" binding:"required"`
}

// AdCategoryClaim is a claim that cointains AdCategory as Data.
type AdCategoryClaim struct {
	Data AdCategory `json:"data" binding:"required"`
	jwt.StandardClaims
}
