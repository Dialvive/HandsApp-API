package models

import "github.com/dgrijalva/jwt-go"

// Email is a simple email.
type Email struct {
	Subject string `json:"_subject"`
	ReplyTo string `json:"_replyTo"`
	Body    string `json:"body"`
}

// EmailInput is a simple email.
type EmailInput struct {
	Subject string `json:"_subject"`
	Body    string `json:"body"`
}

// EmailClaim is a claim that cointains a simple email as Data.
type EmailClaim struct {
	Data Email `json:"data" binding:"required"`
	jwt.StandardClaims
}
