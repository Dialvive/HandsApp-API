package models

import (
	"github.com/dgrijalva/jwt-go"
)

// WordCategory is a category in which a set of words fall into.
type WordCategory struct {
	ID       uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL;primaryKey"`
	NameDe   string `json:"name_de" gorm:"column:name_de;type:TEXT"`
	NameEs   string `json:"name_es" gorm:"column:name_es;type:TEXT"`
	NameEn   string `json:"name_en" gorm:"column:name_en;type:TEXT"`
	NameFr   string `json:"name_fr" gorm:"column:name_fr;type:TEXT"`
	NameIt   string `json:"name_it" gorm:"column:name_it;type:TEXT"`
	NamePt   string `json:"name_pt" gorm:"column:name_pt;type:TEXT"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// WordCategoryInput is a category in which a set of words fall into.
type WordCategoryInput struct {
	NameDe string `json:"name_de"`
	NameEs string `json:"name_es"`
	NameEn string `json:"name_en"`
	NameFr string `json:"name_fr"`
	NameIt string `json:"name_it"`
	NamePt string `json:"name_pt"`
}

// WordCategoryClaim is  a claim that cointains WordCategory as Data.
type WordCategoryClaim struct {
	Data WordCategory `json:"data" binding:"required"`
	jwt.StandardClaims
}
