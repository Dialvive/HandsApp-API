package security

import (
	"API/models"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// KEY is the API key
const KEY string = "SECURITY"

// MinuteTTL is a minute in EPOCH seconds
const MinuteTTL uint64 = 60

// HourTTL is an hour in EPOCH seconds
const HourTTL uint64 = 3600

// DayTTL is a day in EPOCH seconds
const DayTTL uint64 = 86400

// WeekTTL is a day in EPOCH seconds
const WeekTTL uint64 = 604800

// MonthTTL is a month of 30.44 days in EPOCH seconds
const MonthTTL uint64 = 2629743

// YearTTL is a YEAR of 365.24 days in EPOCH seconds
const YearTTL uint64 = 31556926

//GenPassword128 creates a password that is 128 characters long with 32 digits,
// 32 symbols, allowing upper and lower case letters, disallowing repeat characters.
func GenPassword128() (string, error) {
	res1, err := password.Generate(64, 10, 22, false, false)
	if err != nil {
		log.Fatal(err)
	}
	res2, err := password.Generate(64, 10, 22, false, false)
	if err != nil {
		log.Fatal(err)
	}
	res := res1 + res2

	return res, err
}

//GenPassword256 creates a password that is 256 characters long with 64 digits,
// 64 symbols, allowing upper and lower case letters, disallowing repeat characters.
func GenPassword256() (string, error) {
	res1, err := GenPassword128()
	if err != nil {
		log.Fatal(err)
	}
	res2, err := GenPassword128()
	if err != nil {
		log.Fatal(err)
	}
	res := res1 + res2

	return res, err
}

//SecureUint safely processes a string parameter and casts it into a uint
func SecureUint(s string) (uint64, error) {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			err := errors.New("not a valid input")
			return 0, err
		}
	}
	return strconv.ParseUint(s, 10, 64)
}

//SecureString checks for SQL injection strings and characters and returns a
// secure string.
func SecureString(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	str := mysqlEscapeString(strings.ToValidUTF8(s, ""))
	str = "`" + str + "`"
	return str
}

// RemoveBackticks removes every backtick from a given string
func RemoveBackticks(s string) string {
	if s == "" {
		return ""
	}
	replace := map[string]string{
		"`": "",
	}
	for b, a := range replace {
		s = strings.Replace(s, b, a, -1)
	}
	return s
}

// HashPassword uses the Bcrypt hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash, nil
}

// PasswordMatches checks if a passed password matches the original hashed password
func PasswordMatches(hashedPassword, passedPassword string) bool {
	decodedPassword, err := base64.URLEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(decodedPassword, []byte(passedPassword))
	return err == nil
}

// TrimToLength Trims the given string to the given length
func TrimToLength(s string, l int) string {
	if l < 0 {
		l = l * -1
	}
	if len(s) > l {
		return s[0:l]
	}
	return s
}

// CheckKey verifies that a given string matches an API key
func CheckKey(c *gin.Context, s string) bool {
	if s != KEY {
		return false
	}
	return true
}

// RerouteHandler Reroutes every other route to another website
func RerouteHandler(c *gin.Context) {
	c.Redirect(302, "https://handsapp.org")
	return
}

func mysqlEscapeString(s string) string {
	replace := map[string]string{
		"\\":   "",
		"`":    "'",
		"\\0":  "",
		"\n":   "",
		"\r":   "",
		"\x1a": "",
	}
	for b, a := range replace {
		s = strings.Replace(s, b, a, -1)
	}
	return s
}

func CreateJWT(user models.User) string {
	userClaim := models.UserClaim{
		UserName:  RemoveBackticks(user.UserName),
		Mail:      RemoveBackticks(user.Mail),
		Privilege: user.Privilege,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).UTC().Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			NotBefore: time.Now().Add(time.Minute * -5).UTC().Unix(),
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   strconv.Itoa(int(user.ID)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	signedString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		log.Printf("\033[1;31m there has an error creating a jwt for user:\u001B[0m %+v", userClaim)
		log.Println(err.Error())
	}
	return signedString
}

func ParseJWT(tokenFromHeader string, claims *models.UserClaim) error {
	tokenFromHeader = strings.TrimPrefix(tokenFromHeader, "Bearer ")
	_, err := jwt.ParseWithClaims(
		tokenFromHeader,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
		},
	)
	return err
}
