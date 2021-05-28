package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwd, err := security.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	user := models.User{
		FirstName: security.SecureString(input.FirstName),
		LastName:  security.SecureString(input.LastName),
		UserName:  security.SecureString(security.TrimToLength(input.UserName, 30)),
		Mail:      security.SecureString(security.TrimToLength(input.Mail, 252)),
		Password:  pwd,
		Biography: security.SecureString(security.TrimToLength(input.Biography, 140)),
		Mailing:   security.SecureString(security.TrimToLength(input.Mailing, 3)),
		Privilege: security.SecureString(security.TrimToLength(input.Privilege, 3)),
		Points:    uint(input.Points),
		//TODO: TRANSACTION LOCK FOR CREDIT CHANGE
		Credits:  uint(input.Credits),
		LocaleID: uint(input.LocaleID),
		Modified: t,
	}

	if create := models.DB.Create(&user); create.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": create.Error})
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

// PutUser updates a User
func PutUser(c *gin.Context) {

	// Get model if exist
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

	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !security.PasswordMatches(user.Password, input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad password!"})
		return
	}
	pwd, err := security.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&user).Updates(
		models.User{
			ID:        user.ID,
			FirstName: security.SecureString(input.FirstName),
			LastName:  security.SecureString(input.LastName),
			UserName:  security.SecureString(input.UserName),
			Mail:      security.SecureString(input.Mail),
			Password:  pwd,
			Biography: security.SecureString(input.Biography),
			Mailing:   security.SecureString(input.Mailing), Privilege: input.Privilege,
			Points:   uint(input.Points),
			Credits:  uint(input.Credits),
			LocaleID: uint(input.LocaleID),
			Modified: t,
		})

	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Password = "" // NEVER SEND PWD DATA

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PatchUser patches a user with the values that are not default nor equal to the existing one
func PatchUser(c *gin.Context) {

	// Get model if exist
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

	var input models.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !security.PasswordMatches(user.Password, input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad password!"})
		return
	}
	pwd, err := security.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: ALLOW CHANGING PASSWORDS
	user.Password = pwd

	if input.FirstName != "" && input.FirstName != user.FirstName {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" && input.FirstName != user.LastName {
		user.LastName = input.LastName
	}
	if input.UserName != "" && input.UserName != user.UserName {
		user.UserName = input.UserName
	}
	if input.Mail != "" && input.Mail != user.Mail {
		user.Mail = input.Mail
	}

	if input.Biography != "" && input.Biography != user.Biography {
		user.Biography = input.Biography
	}
	if input.Mailing != "" && input.Mailing != user.Mailing {
		user.Mailing = input.Mailing
	}
	if input.Privilege != "" && input.Privilege != user.Privilege {
		user.Privilege = input.Privilege
	}
	if input.Points != 0 && input.Points != user.Points {
		user.Points = input.Points
	}
	if input.Credits != 0 && input.Credits != user.Credits {
		user.Credits = input.Credits
	}
	if input.LocaleID != 0 && input.LocaleID != user.LocaleID {
		user.LocaleID = input.LocaleID
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&user).Updates(
		models.User{
			ID:        user.ID,
			FirstName: security.SecureString(input.FirstName),
			LastName:  security.SecureString(input.LastName),
			UserName:  security.SecureString(security.TrimToLength(input.UserName, 30)),
			Mail:      security.SecureString(security.TrimToLength(input.Mail, 252)),
			Password:  input.Password,
			Biography: security.SecureString(security.TrimToLength(input.Biography, 140)),
			Mailing:   security.SecureString(security.TrimToLength(input.Mailing, 3)),
			Privilege: security.SecureString(security.TrimToLength(user.Privilege, 3)),
			Points:    uint(user.Points),
			Credits:   uint(user.Credits),
			LocaleID:  uint(user.LocaleID),
			Modified:  t,
		})

	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Password = "" // NEVER SEND PWD DATA

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	//TODO: ONLY ALLOW WITH CORRECT PASSWORD OR ADMIN PRIVILEGES
	// Get model if exist
	var user models.User
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
