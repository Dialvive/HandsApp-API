package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetFavoritePhrases retrieves all the favoritePhrases from the DB.
func GetFavoritePhrases(c *gin.Context) {
	var favoritePhrases []models.FavoritePhrases
	models.DB.Find(&favoritePhrases)

	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// CreateFavoritePhrases creates a new favoritePhrase.
func CreateFavoritePhrases(c *gin.Context) {
	var input models.CreateFavoritePhrasesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	favoritePhrases := models.FavoritePhrases{
		PhraseID: input.PhraseID,
		UserID:   input.UserID,
		Modified: t,
	}
	models.DB.Create(&favoritePhrases)
	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// FindFavoritePhrases recieves an user id, and returns all of its favorite phrases.
func FindFavoritePhrases(c *gin.Context) {
	var favorites []models.FavoritePhrases
	if err := models.DB.Where("user_ID = ?", c.Param("id")).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrases not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountFavoritePhrasesP recieves a phrase ID, returns the number of users that user has marked it as favorite.
func CountFavoritePhrasesP(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoritePhrases{}).Where("phrase_ID = ?", c.Param("phraseID")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountFavoritePhrasesU recieves a user ID, returns the number of phrases that user has marked it as favorite.
func CountFavoritePhrasesU(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoritePhrases{}).Where("user_ID = ?", c.Param("userID")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// PutFavoritePhrases updates a favoritePhrase
func PutFavoritePhrases(c *gin.Context) {

	// Get model if exist
	var favoritePhrases models.FavoritePhrases
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", c.Param("userID"), c.Param("phraseID")).First(&favoritePhrases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoritePhrases not found!"})
		return
	}

	var input models.CreateFavoritePhrasesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&favoritePhrases).Updates(
		models.FavoritePhrases{
			PhraseID: input.PhraseID,
			UserID:   input.UserID,
			Modified: t,
		})
	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// DeleteFavoritePhrases deletes a favoritePhrase
func DeleteFavoritePhrases(c *gin.Context) {
	// Get model if exist
	var favoritePhrases models.FavoritePhrases
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", c.Param("userID"), c.Param("phrase_ID")).First(&favoritePhrases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&favoritePhrases)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
