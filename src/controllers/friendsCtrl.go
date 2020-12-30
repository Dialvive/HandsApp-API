package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFriends retrieves all the friends from the DB.
func GetFriends(c *gin.Context) {
	var friends []models.Friend
	models.DB.Find(&friends)

	c.JSON(http.StatusOK, gin.H{"data": friends})
}

// CreateFriend creates a new friend.
func CreateFriend(c *gin.Context) {
	var input models.CreateFriendInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	friend := models.Friend{
		User1ID:      input.User1ID,
		User2ID:      input.User2ID,
		FriendshipID: input.FriendshipID,
		Facebook:     input.Facebook,
		Modified:     t,
	}
	models.DB.Create(&friend)

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// FindFriend recieves an id, and returns an specific friend with that id.
func FindFriend(c *gin.Context) {
	var friend models.Friend

	if err := models.DB.Where("id = ?", c.Param("id")).First(&friend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friend not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// PutFriend updates a friend
func PutFriend(c *gin.Context) {

	// Get model if exist
	var friend models.Friend

	if err := models.DB.Where("id = ?", c.Param("id")).First(&friend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friend not found!"})
		return
	}

	var input models.CreateFriendInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&friend).Updates(
		models.Friend{
			User1ID:      input.User1ID,
			User2ID:      input.User2ID,
			FriendshipID: input.FriendshipID,
			Facebook:     input.Facebook,
			Modified:     t,
		})

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// DeleteFriend deletes a friend
func DeleteFriend(c *gin.Context) {
	// Get model if exist
	var friend models.Friend
	if err := models.DB.Where("id = ?", c.Param("id")).First(&friend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&friend)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
