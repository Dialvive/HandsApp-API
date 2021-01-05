package models

import "github.com/dgrijalva/jwt-go"

// Friendship represents the type of friendship two users have.
type Friendship struct {
	ID       uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL;primaryKey"`
	Name     string `json:"name" gorm:"column:name;type:TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateFriendshipInput represents the type of friendship two users have.
type CreateFriendshipInput struct {
	Name string `json:"name" binding:"required"`
}

// FriendshipClaim is a claim that cointains Friendship as Data.
type FriendshipClaim struct {
	Data Friendship `json:"data" binding:"required"`
	jwt.StandardClaims
}
