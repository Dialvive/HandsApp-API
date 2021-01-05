package models

import "github.com/dgrijalva/jwt-go"

// Friend represents a friendship between two users.
type Friend struct {
	User1ID      uint   `json:"user1_ID" gorm:"column:user1_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	User2ID      uint   `json:"user2_ID" gorm:"column:user2_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	FriendshipID uint   `json:"friendship_ID" gorm:"column:friendship_ID;type:TINYINT NOT NULL"`
	Facebook     bool   `json:"facebook" gorm:"column:facebook;type:BOOLEAN"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateFriendInput represents a friendship between two users.
type CreateFriendInput struct {
	User1ID      uint `json:"user1_ID" binding:"required"`
	User2ID      uint `json:"user2_ID" binding:"required"`
	FriendshipID uint `json:"friendship_ID"`
	Facebook     bool `json:"facebook" binding:"required"`
}

// FriendClaim is a claim that cointains Friend as Data.
type FriendClaim struct {
	Data Friend `json:"data" binding:"required"`
	jwt.StandardClaims
}
