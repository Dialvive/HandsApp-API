package models

import "github.com/dgrijalva/jwt-go"

// Friend represents a friendship between two users.
type Friend struct {
	User1ID      uint   `json:"user1_ID" gorm:"type:INT NOT NULL;primaryKey;autoIncrement:false"`
	User2ID      uint   `json:"user2_ID" gorm:"type:INT NOT NULL;primaryKey;autoIncrement:false"`
	FriendshipID uint   `json:"friendship_ID" gorm:"type:TINYINT NOT NULL"`
	Facebook     bool   `json:"facebook" gorm:"BOOLEAN; default:0"`
	Modified     string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateFriendInput represents a friendship between two users.
type CreateFriendInput struct {
	User1ID      uint `json:"user1_ID" binding:"required"`
	User2ID      uint `json:"user2_ID" binding:"required"`
	FriendshipID uint `json:"friendship_ID" binding:"required"`
	Facebook     bool `json:"facebook" binding:"required"`
}

// FriendClaim is a claim that cointains Friend as Data.
type FriendClaim struct {
	Data Friend `json:"data" binding:"required"`
	jwt.StandardClaims
}
