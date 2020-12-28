package models

// User represents a user of Signa Mundi set of services.
type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name" gorm:"not null"`
	Mail      string `json:"mail" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Biography string `json:"biography"`
	Mailing   string `json:"mailing"`
	Privilege uint   `json:"privilege"`
	Points    uint   `json:"points"`
	Credits   uint   `json:"credits"`
	RegionID  uint   `json:"region_id" gorm:"not null"`
	Creation  string `json:"creation"`
}
