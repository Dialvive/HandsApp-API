package models

// Friendship represents the type of friendship two users have.
type Friendship struct {
	ID       uint   `json:"id" gorm:"TINYINT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}
