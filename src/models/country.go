package models

import "github.com/dgrijalva/jwt-go"

// Country represents a real world country.
type Country struct {
	ID           uint   `json:"ID" gorm:"type:TINYINT;primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"type:TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"type:VARCHAR(4) NOT NULL"`
	Modified     string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateCountryInput type for country POST with automatic ID.
type CreateCountryInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// CountryClaim is a claim that cointains Country as Data.
type CountryClaim struct {
	Data Country `json:"data" binding:"required"`
	jwt.StandardClaims
}
