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

// FindFriends recieves an id, and returns as much friends as there are with that id.
func FindFriends(c *gin.Context) {
	var friends []models.Friend

	err1 := models.DB.Where("user1_ID = ?", c.Param("id")).Find(&friends).Error
	err2 := models.DB.Where("user2_ID = ?", c.Param("id")).Find(&friends).Error

	if err1 != nil && err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friends not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": friends})
}

// FindFriend recieves two IDs, and returns an specific friend with that IDs.
func FindFriend(c *gin.Context) {
	var input models.FindFriendsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var friend models.Friend
	if err1 := models.DB.Model(&friend).First(&input).Error; err1 != nil {
		tupni := models.FindFriendsInput{
			User1ID: input.User2ID,
			User2ID: input.User1ID,
		}
		if err2 := models.DB.Model(&friend).First(&tupni).Error; err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Friend not found!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// CountFriends recieves a user ID, returns the number of users that user has as friends.
func CountFriends(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoritePhrase{}).Where("user1_ID = ?", c.Param("id")).Or("user2_ID = ?", c.Param("id")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
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
	if err := models.DB.Where("user1_ID = ?", c.Param("id")).First(&friend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&friend)

	c.JSON(http.StatusOK, gin.H{"data": true})
}