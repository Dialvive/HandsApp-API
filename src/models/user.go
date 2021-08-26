package models

import (
	"reflect"
)

var (
	UserInputFields = getUserInputFields()
)

// getUserInputFields returns all the fields of UpdateUserInput
// This is made for the update user statement in order to select a few fields
func getUserInputFields() []interface{} {
	userInputType := reflect.TypeOf(UpdateUserInput{})
	fields := make([]interface{}, userInputType.NumField())
	for i := range fields {
		fields[i] = userInputType.Field(i).Name
	}
	return fields
}

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
	GoogleSub      string `json:"google_sub,omitempty" gorm:"type:VARCHAR(64)"`
	FacebookSub    string `json:"facebook_sub,omitempty" gorm:"type:VARCHAR(64)"`
	AppleSub       string `json:"apple_sub,omitempty" gorm:"type:VARCHAR(64)"`
	Picture        string `json:"picture" gorm:"type:VARCHAR(128)"`
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
// The struct has pointers to represent not preset fields
// If a user wants to delete his picture, then use a pointer to an empty field.
// But if the field F is nil, then F will be untouched in the DB.
//     emptyStr := ""
//     // this modifies the database
//     UpdateUserInput{Picture: &emptyStr}
//
//     // this doesn't modify the database
//     UpdateUserInput{Picture: &emptyStr}
type UpdateUserInput struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	UserName    *string `json:"user_name"`
	Mail        *string `json:"mail"`
	Password    *string `json:"password"`
	Biography   *string `json:"biography"`
	Mailing     *string `json:"mailing"`
	LocaleID    *uint   `json:"locale_ID"`
	GoogleSub   *string `json:"google_sub"`
	FacebookSub *string `json:"facebook_sub"`
	AppleSub    *string `json:"apple_sub"`
	Picture     *string `json:"picture"`
}

// LoginForm represents typical login structure, but credential can be a username,
// mail or jwt token (for google)
type LoginForm struct {
	Credential string // username or mail
	Password   string
}

type FacebookForm struct {
	AccessToken              string `json:"accessToken,omitempty"`
	UserID                   string `json:"userID,omitempty"`
	ExpiresIn                uint32 `json:"expiresIn"`
	SignedRequest            string `json:"signedRequest"`
	GraphDomain              string `json:"graphDomain"`
	DataAccessExpirationTime uint   `json:"data_access_expiration_time"`
}
