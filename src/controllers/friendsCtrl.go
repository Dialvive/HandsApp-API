package controllers

import (
	"API/models"
	"API/security"
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
		User1ID:      uint(input.User1ID),
		User2ID:      uint(input.User2ID),
		FriendshipID: uint(input.FriendshipID),
		Facebook:     bool(input.Facebook),
		Modified:     t,
	}
	models.DB.Create(&friend)

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// FindFriends recieves an id, and returns as much friends as there are with that id.
func FindFriends(c *gin.Context) {
	var friends []models.Friend
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID1")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	err1 := models.DB.Where("user1_ID = ?", param).Find(&friends).Error
	err2 := models.DB.Where("user2_ID = ?", param).Find(&friends).Error

	if err1 != nil && err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friends not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": friends})
}

// FindFriend recieves two IDs, and returns an specific friend with that IDs.
func FindFriend(c *gin.Context) {
	var friend models.Friend
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("ID1")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("ID2")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user1_ID = ? AND user2_ID = ?", param1, param2).
		First(&friend).Error; err != nil {

		if err := models.DB.Where("user1_ID = ? AND user2_ID = ?", param2, param1).
			First(&friend).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": "FavoritePhrases not found!"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// CountFriends recieves a user ID, returns the number of users that user has as friends.
func CountFriends(c *gin.Context) {
	var count1 int64
	var count2 int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	err1 := models.DB.Model(&models.Friend{}).Where("user1_ID = ?", param).Count(&count1).Error
	err2 := models.DB.Model(&models.Friend{}).Where("user2_ID = ?", param).Count(&count2).Error

	if err1 != nil && err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friends not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": count1 + count2})
}

// PatchFriend updates a friend
func PatchFriend(c *gin.Context) {

	// Get model if exist
	var friend models.Friend
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("ID1")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("ID2")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.Friend{}).Where("user1_ID = ?", param1).
		Or("user2_ID = ?", param2).Save(&friend).Error; err != nil {

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
			User1ID:      uint(input.User1ID),
			User2ID:      uint(input.User2ID),
			FriendshipID: uint(input.FriendshipID),
			Facebook:     bool(input.Facebook),
			Modified:     t,
		})

	c.JSON(http.StatusOK, gin.H{"data": friend})
}

// DeleteFriend deletes a friend
func DeleteFriend(c *gin.Context) {
	// Get model if exist
	var friend models.Friend
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("ID1")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("ID2")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.Friend{}).Where("user1_ID = ?", param1).
		Or("user2_ID = ?", param2).First(&friend).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&friend)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
