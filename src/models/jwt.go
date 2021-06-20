package models

import "github.com/dgrijalva/jwt-go"

// HandsAppJWT is a wrapper that contains information about the jwt token:
type HandsAppJWT struct {
	// the raw jwt from a signed user
	Token string
	// its a random string within the signed Token
	CsrfToken string
	// time in unix timestamp (utc)
	ExpireAt int64
}

// UserClaim is a claim the jwt sent and received by the user but parsed
type UserClaim struct {
	UserName  string `json:"user_name"`
	Mail      string `json:"mail"`
	Privilege string `json:"privilege"`
	CsrfToken string
	jwt.StandardClaims
}
