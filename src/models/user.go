package models

import "github.com/dgrijalva/jwt-go"

// User represents a user of Signa Mundi set of services.
type User struct {
	ID             uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	FirstName      string `json:"first_name" gorm:"column:first_name;type:TEXT"`
	LastName       string `json:"last_name" gorm:"column:last_name;type:TEXT"`
	UserName       string `json:"user_name" gorm:"column:user_name;type:VARCHAR(32) NOT NULL"`
	Mail           string `json:"mail" gorm:"column:mail;type:VARCHAR(254) NOT NULL"`
	Password       string `json:"password" gorm:"column:password;type:TEXT NOT NULL"`
	Biography      string `json:"biography" gorm:"column:biography;type:TEXT"`
	Mailing        string `json:"mailing" gorm:"type:set('notification', 'promotion', 'advertising');default:'notification'"`
	Privilege      string `json:"privilege" gorm:"type:enum('child', 'adult', 'editor', 'super user');default:'adult'"`
	Points         uint   `json:"points" gorm:"column:points;type:INT; default:0"`
	Credits        uint   `json:"credits" gorm:"column:credits;type:INT; default:0"`
	LocaleID       uint   `json:"locale_id" gorm:"column:locale_ID;type:INT NOT NULL"`
	Modified       string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
	GoogleSub      string `json:"google_sub,omitempty"`
	FacebookSub    string `json:"facebook_sub,omitempty"`
	AppleSub       string `json:"apple_sub,omitempty"`
	Picture        string `json:"picture"`
	SubscriberType string `json:"subscriber_type" gorm:"type:enum('free', 'premium');default:'free'"`
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
	Privilege string `json:"privilege" binding:"required"`
	Points    uint   `json:"points"`
	Credits   uint   `json:"credits"`
	LocaleID  uint   `json:"locale_ID" binding:"required"`
}

// UpdateUserInput represents a user of Signa Mundi set of services.
type UpdateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Biography string `json:"biography"`
	Mailing   string `json:"mailing"`
	Privilege string `json:"privilege"`
	Points    uint   `json:"points"`
	Credits   uint   `json:"credits"`
	LocaleID  uint   `json:"locale_ID"`
}

// UserClaim is  a claim that cointains User as Data.
type UserClaim struct {
	UserName  string `json:"user_name"`
	Mail      string `json:"mail"`
	Privilege string `json:"privilege"`
	jwt.StandardClaims
}

// LoginForm represents typical login structure, but credential can be a username,
// mail or jwt token (for google)
type LoginForm struct {
	Credential string // username or mail
	Password   string
}
