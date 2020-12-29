package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all the users from the DB.
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserName:  input.UserName,
		Mail:      input.Mail,
		Password:  input.Password,
		Biography: input.Biography,
		Mailing:   input.Mailing,
		Privilege: input.Privilege,
		Points:    input.Points,
		Credits:   input.Credits,
		RegionID:  input.RegionID,
		Creation:  t,
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// FindUser recieves an id, and returns an specific user with that id.
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PutUser updates a User
func PutUser(c *gin.Context) {

	// Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&user).Updates(
		models.User{
			ID:        user.ID,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			UserName:  input.UserName,
			Mail:      input.Mail,
			Password:  input.Password,
			Biography: input.Biography,
			Mailing:   input.Mailing,
			Privilege: input.Privilege,
			Points:    input.Points,
			Credits:   input.Credits,
			RegionID:  input.RegionID,
			Creation:  t,
		})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PatchUser patches a user with the values that are not default nor equal to the existing one
func PatchUser(c *gin.Context) {

	// Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.PatchUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	if input.Password != "" && input.Password != user.Password {
		user.Password = input.Password
	}
	if input.Biography != "" && input.Biography != user.Biography {
		user.Biography = input.Biography
	}
	if input.Mailing != "" && input.Mailing != user.Mailing {
		user.Mailing = input.Mailing
	}
	if input.Privilege != 0 && input.Privilege != user.Privilege {
		user.Privilege = input.Privilege
	}
	if input.Points != 0 && input.Points != user.Points {
		user.Points = input.Points
	}
	if input.Credits != 0 && input.Credits != user.Credits {
		user.Credits = input.Credits
	}
	if input.RegionID != 0 && input.RegionID != user.RegionID {
		user.RegionID = input.RegionID
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&user).Updates(
		models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserName:  user.UserName,
			Mail:      user.Mail,
			Password:  user.Password,
			Biography: user.Biography,
			Mailing:   user.Mailing,
			Privilege: user.Privilege,
			Points:    user.Points,
			Credits:   user.Credits,
			RegionID:  user.RegionID,
			Creation:  t,
		})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
