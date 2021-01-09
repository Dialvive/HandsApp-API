package controllers

import (
	"API/models"
	"API/security"
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

// CreateFavoritePhrases creates a new favoritePhrase.
func CreateFavoritePhrases(c *gin.Context) {
	var input models.FavoritePhraseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	favoritePhrases := models.FavoritePhrase{
		PhraseID: uint(input.PhraseID),
		UserID:   uint(input.UserID),
		Modified: t,
	}
	models.DB.Create(&favoritePhrases)
	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// FindFavoritePhrases recieves an user id, and returns all of its favorite phrases.
func FindFavoritePhrases(c *gin.Context) {
	var favorites []models.FavoritePhrase
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("userID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user_ID = ?", param).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrases not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountFavoritePhrasesP recieves a phrase ID, returns the number of users that user has
// marked it as favorite.
func CountFavoritePhrasesP(c *gin.Context) {
	var count int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("phraseID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.FavoritePhrase{}).Where("phrase_ID = ?", param).
		Count(&count).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountFavoritePhrasesU recieves a user ID, returns the number of phrases that user has
// marked it as favorite.
func CountFavoritePhrasesU(c *gin.Context) {
	var count int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("userID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.FavoritePhrase{}).Where("user_ID = ?", param).
		Count(&count).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// PutFavoritePhrases updates a favoritePhrase
func PutFavoritePhrases(c *gin.Context) {

	// Get model if exist
	var favoritePhrases models.FavoritePhrase
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("userID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("phraseID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", param1, param2).
		First(&favoritePhrases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "FavoritePhrases not found!"})
		return
	}

	var input models.FavoritePhraseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&favoritePhrases).Updates(
		models.FavoritePhrase{
			PhraseID: uint(input.PhraseID),
			UserID:   uint(input.UserID),
			Modified: t,
		})
	c.JSON(http.StatusOK, gin.H{"data": favoritePhrases})
}

// DeleteFavoritePhrases deletes a favoritePhrase
func DeleteFavoritePhrases(c *gin.Context) {
	// Get model if exist
	var favoritePhrases models.FavoritePhrase
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("userID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("phraseID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user_ID = ? AND phrase_ID = ?", param1, param2).
		First(&favoritePhrases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&favoritePhrases)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
