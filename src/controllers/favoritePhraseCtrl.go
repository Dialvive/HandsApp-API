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

// FindFavoritePhrases recieves an user id, and returns all of its favorite phrases.
func FindFavoritePhrases(c *gin.Context) {
	var favorites []models.FavoritePhrase
	if err := models.DB.Where("user_ID = ?", c.Param("id")).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrases not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountFavoritePhrasesP recieves a phrase ID, returns the number of users that user has marked it as favorite.
func CountFavoritePhrasesP(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoritePhrase{}).Where("phrase_ID = ?", c.Param("phraseId")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountFavoritePhrasesU recieves a user ID, returns the number of phrases that user has marked it as favorite.
func CountFavoritePhrasesU(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoritePhrase{}).Where("user_ID = ?", c.Param("userId")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// PutFavoritePhrase updates a favoritePhrase
func PutFavoritePhrase(c *gin.Context) {

	// Get model if exist
	var favoritePhrase models.FavoritePhrase
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", c.Param("userId"), c.Param("phraseId")).First(&favoritePhrase).Error; err != nil {
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
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", c.Param("userId"), c.Param("phrase_Id")).First(&favoritePhrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&favoritePhrase)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
