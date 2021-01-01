package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFavoritePhrases retrieves all the favoritePhrases from the DB.
func GetFavoritePhrases(c *gin.Context) {
	var favoritePhrases []models.FavoritePhrase
	models.DB.Find(&favoritePhrases)

	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// CreateFavoritePhrase creates a new favoritePhrase.
func CreateFavoritePhrase(c *gin.Context) {
	var input models.CreateFavoritePhraseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	favoritePhrase := models.FavoritePhrase{
		PhraseID: input.PhraseID,
		UserID:   input.UserID,
		Modified: t,
	}
	models.DB.Create(&favoritePhrase)

	c.JSON(http.StatusOK, gin.H{"data": favoritePhrase})
}

// FindFavoritePhrase recieves an id, and returns an specific favoritePhrase with that id.
func FindFavoritePhrase(c *gin.Context) {
	var favoritePhrase models.FavoritePhrase

	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoritePhrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoritePhrase not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favoritePhrase})
}

// PutFavoritePhrase updates a favoritePhrase
func PutFavoritePhrase(c *gin.Context) {

	// Get model if exist
	var favoritePhrase models.FavoritePhrase

	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoritePhrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoritePhrase not found!"})
		return
	}

	var input models.CreateFavoritePhraseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&favoritePhrase).Updates(
		models.FavoritePhrase{
			PhraseID: input.PhraseID,
			UserID:   input.UserID,
			Modified: t,
		})

	c.JSON(http.StatusOK, gin.H{"data": favoritePhrase})
}

// DeleteFavoritePhrase deletes a favoritePhrase
func DeleteFavoritePhrase(c *gin.Context) {
	// Get model if exist
	var favoritePhrase models.FavoritePhrase
	if err := models.DB.Where("id = ?", c.Param("id")).First(&favoritePhrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&favoritePhrase)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
