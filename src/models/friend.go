package models

// Friend represents a friendship between two users.
type Friend struct {
	User1ID      uint   `json:"user1_ID" gorm:"INT NOT NULL"`
	User2ID      uint   `json:"user2_ID" gorm:"INT NOT NULL"`
	FriendshipID uint   `json:"friendship_ID" gorm:"TINYINT NOT NULL"`
	Facebook     bool   `json:"facebook" gorm:"BOOLEAN; default:0"`
	Modified     string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateFriendInput represents a friendship between two users.
type CreateFriendInput struct {
	User1ID      uint `json:"user1_ID" binding:"required"`
	User2ID      uint `json:"user2_ID" binding:"required"`
	FriendshipID uint `json:"friendship_ID" binding:"required"`
	Facebook     bool `json:"facebook" binding:"required"`
}

// FindFriendsInput represents a friendship between two users.
type FindFriendsInput struct {
	User1ID uint `json:"user1_ID" binding:"required"`
	User2ID uint `json:"user2_ID" binding:"required"`
}
