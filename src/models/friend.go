package models

// Friend represents a friendship between two users.
type Friend struct {
	User1ID      uint   `json:"user1_ID" gorm:"not null"`
	User2ID      uint   `json:"user2_ID" gorm:"not null"`
	FriendshipID string `json:"friendship_ID" gorm:"not null"`
	Facebook     bool   `json:"facebook"`
	Creation     string `json:"creation"`
}
