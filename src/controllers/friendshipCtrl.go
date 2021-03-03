package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFriendships retrieves all the friendships from the DB.
func GetFriendships(c *gin.Context) {
	var friendships []models.Friendship
	models.DB.Find(&friendships)

	for i := range friendships {
		friendships[i].NameDe = security.RemoveBackticks(friendships[i].NameDe)
		friendships[i].NameEs = security.RemoveBackticks(friendships[i].NameEs)
		friendships[i].NameEn = security.RemoveBackticks(friendships[i].NameEn)
		friendships[i].NameFr = security.RemoveBackticks(friendships[i].NameFr)
		friendships[i].NameIt = security.RemoveBackticks(friendships[i].NameIt)
		friendships[i].NamePt = security.RemoveBackticks(friendships[i].NamePt)
	}

	c.JSON(http.StatusOK, gin.H{"data": friendships})
}

// CreateFriendship creates a new friendship.
func CreateFriendship(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.FriendshipInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	friendship := models.Friendship{
		NameDe:   security.SecureString(input.NameDe),
		NameEs:   security.SecureString(input.NameEs),
		NameEn:   security.SecureString(input.NameEn),
		NameFr:   security.SecureString(input.NameFr),
		NameIt:   security.SecureString(input.NameIt),
		NamePt:   security.SecureString(input.NamePt),
		Modified: t}
	models.DB.Create(&friendship)

	friendship.NameDe = security.RemoveBackticks(friendship.NameDe)
	friendship.NameEs = security.RemoveBackticks(friendship.NameEs)
	friendship.NameEn = security.RemoveBackticks(friendship.NameEn)
	friendship.NameFr = security.RemoveBackticks(friendship.NameFr)
	friendship.NameIt = security.RemoveBackticks(friendship.NameIt)
	friendship.NamePt = security.RemoveBackticks(friendship.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// FindFriendship recieves an id, and returns an specific friendship with that id.
func FindFriendship(c *gin.Context) {
	var friendship models.Friendship
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friendship not found!"})
		return
	}

	friendship.NameDe = security.RemoveBackticks(friendship.NameDe)
	friendship.NameEs = security.RemoveBackticks(friendship.NameEs)
	friendship.NameEn = security.RemoveBackticks(friendship.NameEn)
	friendship.NameFr = security.RemoveBackticks(friendship.NameFr)
	friendship.NameIt = security.RemoveBackticks(friendship.NameIt)
	friendship.NamePt = security.RemoveBackticks(friendship.NamePt)
	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// PatchFriendship updates a friendship
func PatchFriendship(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var friendship models.Friendship
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friendship not found!"})
		return
	}

	var input models.FriendshipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&friendship).Updates(
		models.Friendship{
			ID:       friendship.ID,
			NameDe:   security.SecureString(input.NameDe),
			NameEs:   security.SecureString(input.NameEs),
			NameEn:   security.SecureString(input.NameEn),
			NameFr:   security.SecureString(input.NameFr),
			NameIt:   security.SecureString(input.NameIt),
			NamePt:   security.SecureString(input.NamePt),
			Modified: t,
		})

	friendship.NameDe = security.RemoveBackticks(friendship.NameDe)
	friendship.NameEs = security.RemoveBackticks(friendship.NameEs)
	friendship.NameEn = security.RemoveBackticks(friendship.NameEn)
	friendship.NameFr = security.RemoveBackticks(friendship.NameFr)
	friendship.NameIt = security.RemoveBackticks(friendship.NameIt)
	friendship.NamePt = security.RemoveBackticks(friendship.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": friendship})
}

// DeleteFriendship deletes a friendship
func DeleteFriendship(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var friendship models.Friendship
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&friendship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&friendship)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
