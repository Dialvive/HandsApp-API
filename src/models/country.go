package models

import "github.com/dgrijalva/jwt-go"

// Country represents a real world country.
type Country struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL;primaryKey;"`
	NameDe       string `json:"name_de" gorm:"column:name_de;type:TEXT"`
	NameEs       string `json:"name_es" gorm:"column:name_es;type:TEXT"`
	NameEn       string `json:"name_en" gorm:"column:name_en;type:TEXT"`
	NameFr       string `json:"name_fr" gorm:"column:name_fr;type:TEXT"`
	NameIt       string `json:"name_it" gorm:"column:name_it;type:TEXT"`
	NamePt       string `json:"name_pt" gorm:"column:name_pt;type:TEXT"`
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;type:VARCHAR(4) NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateCountryInput type for country POST with automatic ID.
type CreateCountryInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// UpdateCountryInput type for country POST with automatic ID.
type UpdateCountryInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation"`
}

// CountryClaim is a claim that cointains Country as Data.
type CountryClaim struct {
	Data Country `json:"data" binding:"required"`
	jwt.StandardClaims
}
