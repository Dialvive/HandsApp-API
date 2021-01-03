package models

import "github.com/dgrijalva/jwt-go"

// User represents a user of Signa Mundi set of services.
type User struct {
	ID        uint   `json:"id" gorm:"type:INT; primaryKey"`
	FirstName string `json:"first_name" gorm:"type:TEXT"`
	LastName  string `json:"last_name" gorm:"type:TEXT"`
	UserName  string `json:"user_name" gorm:"type:VARCHAR(32) NOT NULL"`
	Mail      string `json:"mail" gorm:"type:VARCHAR(254) NOT NULL"`
	Password  string `json:"password" gorm:"type:TEXT NOT NULL"`
	Biography string `json:"biography" gorm:"type:TEXT"`
	Mailing   string `json:"mailing" gorm:"type:VARCHAR(3)"`
	Privilege uint   `json:"privilege" gorm:"type:VARCHAR(3) NOT NULL"`
	Points    uint   `json:"points" gorm:"type:INT; default:0"`
	Credits   uint   `json:"credits" gorm:"type:INT; default:0"`
	RegionID  uint   `json:"region_id" gorm:"type:INT NOT NULL"`
	Modified  string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateUserInput represents a user of Signa Mundi set of services.
type CreateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Biography string `json:"biography"`
	Mailing   string `json:"mailing" binding:"required"`
	Privilege uint   `json:"privilege" binding:"required"`
	Points    uint   `json:"points"`
	Credits   uint   `json:"credits"`
	RegionID  uint   `json:"region_ID" binding:"required"`
}

// PatchUserInput represents a user of Signa Mundi set of services.
type PatchUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Biography string `json:"biography"`
	Mailing   string `json:"mailing"`
	Privilege uint   `json:"privilege"`
	Points    uint   `json:"points"`
	Credits   uint   `json:"credits"`
	RegionID  uint   `json:"region_ID"`
}

// UserClaim is  a claim that cointains User as Data.
type UserClaim struct {
	Data User `json:"data" binding:"required"`
	jwt.StandardClaims
}
