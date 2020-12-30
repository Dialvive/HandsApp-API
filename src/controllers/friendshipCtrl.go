package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFriendships retrieves all the friendships from the DB.
func GetFriendships(c *gin.Context) {
	var friendships []models.Friendship
	models.DB.Find(&friendships)

	c.JSON(http.StatusOK, gin.H{"data": friendships})
}

// CreateFriendship creates a new friendship.
func CreateFriendship(c *gin.Context) {
	var input models.CreateFriendshipInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	friendship := models.Friendship{
		Name:     input.Name,
		Modified: t}
	models.DB.Create(&friendship)

	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// FindFriendship recieves an id, and returns an specific friendship with that id.
func FindFriendship(c *gin.Context) {
	var friendship models.Friendship

	if err := models.DB.Where("id = ?", c.Param("id")).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friendship not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// PatchFriendship updates a friendship
func PatchFriendship(c *gin.Context) {

	// Get model if exist
	var friendship models.Friendship

	if err := models.DB.Where("id = ?", c.Param("id")).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friendship not found!"})
		return
	}

	var input models.CreateFriendshipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&friendship).Updates(
		models.Friendship{
			ID:       friendship.ID,
			Name:     input.Name,
			Modified: t,
		})

	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// DeleteFriendship deletes a friendship
func DeleteFriendship(c *gin.Context) {
	// Get model if exist
	var friendship models.Friendship
	if err := models.DB.Where("id = ?", c.Param("id")).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&friendship)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
