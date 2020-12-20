package models

// User represents a user of Signa Mundi set of services.
type User struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	FirstName      string `json:"name" gorm:"not null"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name" gorm:"not null"`
	Mail           string `json:"mail" gorm:"not null"`
	Password       string `json:"password" gorm:"not null"`
	Biography      string `json:"biography"`
	ProfilePicture string `json:"profile_picture"`
	Mailing        string `json:"mailing"`
	Type           string `json:"type"`
	Privilege      uint   `json:"privilege"`
	Country        uint   `json:"country" gorm:"not null"`
}
