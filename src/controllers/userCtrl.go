package controllers

import (
	"API/models"
	"API/security"
	services "API/services/users"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/huandu/facebook/v2"
	"google.golang.org/api/idtoken"
	"net/http"
	"os"
)

var (
	userService = services.UserService{}
	facebookApp = facebook.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_APP_SECRET"))
)

const (
	HandsAppCsrfToken = "HandsApp-Csrf-Token"
	HandsAppSession   = "__Host-ha-session"
	GinKeyUser        = "GinKeyUser"
)

func init() {
	facebookApp.EnableAppsecretProof = true
}

// GetUsers retrieves all the users from the DB.
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	for i := range users {
		users[i].Biography = security.RemoveBackticks(users[i].Biography)
		users[i].FirstName = security.RemoveBackticks(users[i].FirstName)
		users[i].LastName = security.RemoveBackticks(users[i].LastName)
		users[i].UserName = security.RemoveBackticks(users[i].UserName)
		users[i].Mail = security.RemoveBackticks(users[i].Mail)
		users[i].Mailing = security.RemoveBackticks(users[i].Mailing)
		users[i].Password = "" // NEVER SEND PWD DATA
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, dbError := userService.SignWithHandsApp(input)

	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbError.Error()})
		return
	}
	putCsrfToken(c, user)
	c.JSON(http.StatusOK, gin.H{"data": user.ExpireAt})
}

func Login(c *gin.Context) {
	var form models.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	token, err := userService.Login(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	putCsrfToken(c, token)
	c.JSON(http.StatusOK, gin.H{"data": token.ExpireAt})
}

// CreateUserWithGoogle with a token, there is no need for a password
func CreateUserWithGoogle(c *gin.Context) {
	var form models.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload, err := idtoken.Validate(context.Background(), form.Credential, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims := payload.Claims
	user := models.User{
		LocaleID:  1,
		Mail:      claims["email"].(string),
		UserName:  payload.Subject,
		FirstName: claims["given_name"].(string),
		LastName:  claims["family_name"].(string),
		GoogleSub: payload.Subject,
		Picture:   claims["picture"].(string),
	}
	userSaved, err := userService.SignWithGoogle(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	putCsrfToken(c, userSaved)
	c.JSON(http.StatusOK, gin.H{"data": userSaved.ExpireAt})
}

func CreateUserWithFacebook(c *gin.Context) {
	var form models.FacebookForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := facebookApp.Session(form.AccessToken)

	if sessionError := session.Validate(); sessionError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": sessionError})
		return
	}

	facebookInfo, err := session.Get("/me", facebook.Params{
		"fields": "first_name,last_name,email,picture{url}",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	picture, _ := facebookInfo.Get("picture.data.url").(string)
	mail, _ := facebookInfo["email"].(string)

	user := models.User{
		LocaleID:    1,
		Mail:        mail,
		UserName:    facebookInfo["id"].(string),
		FirstName:   facebookInfo["first_name"].(string),
		LastName:    facebookInfo["last_name"].(string),
		FacebookSub: facebookInfo["id"].(string),
		Picture:     picture,
	}
	userToken, err := userService.SignWithFacebook(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	putCsrfToken(c, userToken)
	c.JSON(http.StatusOK, gin.H{"data": userToken.ExpireAt})
}

func putCsrfToken(c *gin.Context, userToken models.HandsAppJWT) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(HandsAppSession, userToken.Token, int(security.TokenLifetime.Seconds()), "", "", true, true)
	c.Writer.Header().Set(HandsAppCsrfToken, userToken.CsrfToken)
}

// FindUser recieves an id, and returns an specific user with that id.
func FindUser(c *gin.Context) {
	var user models.User
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Password = "" // NEVER SEND PWD DATA

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PatchUser patches a user with the values of models.UpdateUserInput that are not nil
func PatchUser(c *gin.Context) {
	id, err := security.SecureUint(c.Param("ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	jwtClaims := c.MustGet(GinKeyUser).(models.UserClaim)

	if jwtClaims.Subject != c.Param("ID") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can't edit this user"})
		return
	}

	var input models.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.Update(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user = services.Unsafe(user)
	user.Password = "" // NEVER SEND PWD DATA

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	jwtClaims := c.MustGet(GinKeyUser).(models.UserClaim)

	isRoot := jwtClaims.Privilege == "super user"
	if !isRoot && jwtClaims.Subject != c.Param("ID") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can't delete this user"})
		return
	}

	param, err := security.SecureUint(c.Param("ID"))
	if err != nil || param == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid id"})
		return
	}

	tx := models.DB.Delete(&models.User{ID: uint(param)})
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": tx.Error})
		return
	}

	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "user doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
