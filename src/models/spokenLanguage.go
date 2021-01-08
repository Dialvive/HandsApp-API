package models

import "github.com/dgrijalva/jwt-go"

// SpokenLanguage represents a language such as English, Spanish, etc.
type SpokenLanguage struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL; primaryKey"`
	NameDe       string `json:"name_de" gorm:"column:name_de;type:TEXT"`
	NameEs       string `json:"name_es" gorm:"column:name_es;type:TEXT"`
	NameEn       string `json:"name_en" gorm:"column:name_en;type:TEXT"`
	NameFr       string `json:"name_fr" gorm:"column:name_fr;type:TEXT"`
	NameIt       string `json:"name_it" gorm:"column:name_it;type:TEXT"`
	NamePt       string `json:"name_pt" gorm:"column:name_pt;type:TEXT"`
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;type:VARCHAR(4) NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateSpokenLanguageInput type with automatic ID.
type CreateSpokenLanguageInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// UpdateSpokenLanguageInput type with automatic ID.
type UpdateSpokenLanguageInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation"`
}

// SpokenLanguageClaim is  a claim that cointains SpokenLanguage as Data.
type SpokenLanguageClaim struct {
	Data SpokenLanguage `json:"data" binding:"required"`
	jwt.StandardClaims
}
