package models

// User represents a user of Signa Mundi set of services.
type User struct {
	ID        uint   `json:"id" gorm:"INT; primaryKey"`
	FirstName string `json:"first_name" gorm:"TEXT"`
	LastName  string `json:"last_name" gorm:"TEXT"`
	UserName  string `json:"user_name" gorm:"VARCHAR(32) NOT NULL"`
	Mail      string `json:"mail" gorm:"VARCHAR(254) NOT NULL"`
	Password  string `json:"password" gorm:"TEXT NOT NULL"`
	Biography string `json:"biography" gorm:"TEXT"`
	Mailing   string `json:"mailing" gorm:"VARCHAR(3)"`
	Privilege uint   `json:"privilege" gorm:"VARCHAR(3) NOT NULL"`
	Points    uint   `json:"points" gorm:"INT; default:0"`
	Credits   uint   `json:"credits" gorm:"INT; default:0"`
	RegionID  uint   `json:"region_id" gorm:"INT NOT NULL"`
	Creation  string `json:"creation" gorm:"TIMESTAMP"`
}
