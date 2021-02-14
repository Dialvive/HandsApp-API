package models

import "github.com/dgrijalva/jwt-go"

// AdCategory represents the type of ad an ad is.
type AdCategory struct {
	ID       uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL;primaryKey"`
	NameDe   string `json:"name_de" gorm:"column:name_de;type:TEXT"`
	NameEs   string `json:"name_es" gorm:"column:name_es;type:TEXT"`
	NameEn   string `json:"name_en" gorm:"column:name_en;type:TEXT"`
	NameFr   string `json:"name_fr" gorm:"column:name_fr;type:TEXT"`
	NameIt   string `json:"name_it" gorm:"column:name_it;type:TEXT"`
	NamePt   string `json:"name_pt" gorm:"column:name_pt;type:TEXT"`
	Cost     uint   `json:"cost" gorm:"column:cost;type:INT NOT NULL"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateAdCategoryInput type for ad_category POST with automatic ID.
type CreateAdCategoryInput struct {
	NameDe string `json:"name_de"`
	NameEs string `json:"name_es"`
	NameEn string `json:"name_en"`
	NameFr string `json:"name_fr"`
	NameIt string `json:"name_it"`
	NamePt string `json:"name_pt"`
	Cost   uint   `json:"cost" binding:"required"`
}

// UpdateAdCategoryInput type for ad_category POST with automatic ID.
type UpdateAdCategoryInput struct {
	NameDe string `json:"name_de"`
	NameEs string `json:"name_es"`
	NameEn string `json:"name_en"`
	NameFr string `json:"name_fr"`
	NameIt string `json:"name_it"`
	NamePt string `json:"name_pt"`
	Cost   uint   `json:"cost"`
}

// AdCategoryClaim is a claim that cointains AdCategory as Data.
type AdCategoryClaim struct {
	Data AdCategory `json:"data" binding:"required"`
	jwt.StandardClaims
}
