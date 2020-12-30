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
	Modified  string `json:"modified" gorm:"TIMESTAMP"`
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
