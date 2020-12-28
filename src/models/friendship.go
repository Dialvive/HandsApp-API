package models

// Friendship represents the type of friendship two users have.
type Friendship struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Creation string `json:"creation"`
}
