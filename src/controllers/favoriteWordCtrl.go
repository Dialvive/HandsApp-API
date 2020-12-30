package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFavoriteWords retrieves all the favoriteWords from the DB.
func GetFavoriteWords(c *gin.Context) {
	var favoriteWords []models.FavoriteWord
	models.DB.Find(&favoriteWords)

	c.JSON(http.StatusOK, gin.H{"data": favoriteWords})
}

// CreateFavoriteWord creates a new favoriteWord.
func CreateFavoriteWord(c *gin.Context) {
	var input models.CreateFavoriteWordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	favoriteWord := models.FavoriteWord{
		WordID: input.WordID,
		UserID: input.UserID,
	}
	models.DB.Create(&favoriteWord)

	c.JSON(http.StatusOK, gin.H{"data": favoriteWord})
}

// FindFavoriteWord recieves an id, and returns an specific favoriteWord with that id.
func FindFavoriteWord(c *gin.Context) {
	var favoriteWord models.FavoriteWord

	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoriteWord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoriteWord not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favoriteWord})
}

// PatchFavoriteWord updates a favoriteWord
func PatchFavoriteWord(c *gin.Context) {

	// Get model if exist
	var favoriteWord models.FavoriteWord

	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoriteWord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoriteWord not found!"})
		return
	}

	var input models.CreateFavoriteWordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&favoriteWord).Updates(
		models.FavoriteWord{
			WordID:   input.WordID,
			UserID:   input.UserID,
			Modified: t,
		})

	c.JSON(http.StatusOK, gin.H{"data": favoriteWord})
}

// DeleteFavoriteWord deletes a favoriteWord
func DeleteFavoriteWord(c *gin.Context) {
	// Get model if exist
	var favoriteWord models.FavoriteWord
	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoriteWord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&favoriteWord)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
