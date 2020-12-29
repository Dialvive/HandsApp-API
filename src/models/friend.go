package models

// Friend represents a friendship between two users.
type Friend struct {
	User1ID      uint   `json:"user1_ID" gorm:"INT NOT NULL"`
	User2ID      uint   `json:"user2_ID" gorm:"INT NOT NULL"`
	FriendshipID uint   `json:"friendship_ID" gorm:"TINYINT NOT NULL"`
	Facebook     bool   `json:"facebook" gorm:"BOOLEAN; default:0"`
	Creation     string `json:"creation" gorm:"TIMESTAMP"`
}
