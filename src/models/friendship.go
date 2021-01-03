package models

import "github.com/dgrijalva/jwt-go"

// Friendship represents the type of friendship two users have.
type Friendship struct {
	ID       uint   `json:"ID" gorm:"type:TINYINT; primaryKey"`
	Name     string `json:"name" gorm:"type:TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateFriendshipInput represents the type of friendship two users have.
type CreateFriendshipInput struct {
	Name string `json:"cost" binding:"required"`
}

// FriendshipClaim is a claim that cointains Friendship as Data.
type FriendshipClaim struct {
	Data Friendship `json:"data" binding:"required"`
	jwt.StandardClaims
}
