package models

import "github.com/dgrijalva/jwt-go"

// Region represents an administrative division of a country.
type Region struct {
	ID        uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	Name      string `json:"name" gorm:"column:name;type:TEXT NOT NULL"`
	CountryID uint   `json:"country_ID" gorm:"column:country_ID;type:TINYINT NOT NULL"`
	Modified  string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateRegionInput represents an administrative division of a country.
type CreateRegionInput struct {
	Name      string `json:"name" binding:"required"`
	CountryID int    `json:"country_ID" binding:"required"`
}

// UpdateRegionInput represents an administrative division of a country.
type UpdateRegionInput struct {
	Name      string `json:"name"`
	CountryID int    `json:"country_ID"`
}

// RegionClaim is  a claim that cointains Region as Data.
type RegionClaim struct {
	Data Region `json:"data" binding:"required"`
	jwt.StandardClaims
}
