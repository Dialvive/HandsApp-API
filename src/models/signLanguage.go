package models

import "github.com/dgrijalva/jwt-go"

// SignLanguage represents a Sign Language such as Mexican Sign Language.
type SignLanguage struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:SMALLINT NOT NULL;primaryKey"`
	NameDe       string `json:"name_de" gorm:"column:name_de;type:TEXT"`
	NameEs       string `json:"name_es" gorm:"column:name_es;type:TEXT"`
	NameEn       string `json:"name_en" gorm:"column:name_en;type:TEXT"`
	NameFr       string `json:"name_fr" gorm:"column:name_fr;type:TEXT"`
	NameIt       string `json:"name_it" gorm:"column:name_it;type:TEXT"`
	NamePt       string `json:"name_pt" gorm:"column:name_pt;type:TEXT"`
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;type:VARCHAR(8) NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateSignLanguageInput type with automatic ID.
type CreateSignLanguageInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// UpdateSignLanguageInput type with automatic ID.
type UpdateSignLanguageInput struct {
	NameDe       string `json:"name_de"`
	NameEs       string `json:"name_es"`
	NameEn       string `json:"name_en"`
	NameFr       string `json:"name_fr"`
	NameIt       string `json:"name_it"`
	NamePt       string `json:"name_pt"`
	Abbreviation string `json:"abbreviation"`
}

// SignLanguageClaim is  a claim that cointains SignLanguage as Data.
type SignLanguageClaim struct {
	Data SignLanguage `json:"data" binding:"required"`
	jwt.StandardClaims
}
