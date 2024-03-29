package controllers

import (
	"API/models"
	"API/security"
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

// CreateFavoriteWords creates a new favoriteWords.
func CreateFavoriteWords(c *gin.Context) {
	var input models.FavoriteWordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	favoriteWords := models.FavoriteWord{
		WordID:   uint(input.WordID),
		UserID:   uint(input.UserID),
		Modified: t,
	}
	models.DB.Create(&favoriteWords)

	c.JSON(http.StatusOK, gin.H{"data": favoriteWords})
}

// FindFavoriteWords recieves an user id, and returns all of its favorite words.
func FindFavoriteWords(c *gin.Context) {
	var favorites []models.FavoriteWord
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("userID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user_ID = ?", param).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Words not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountFavoriteWordsP recieves a word ID, returns the number of users that user has marked
// it as favorite.
func CountFavoriteWordsP(c *gin.Context) {
	var count int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("wordID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.FavoriteWord{}).
		Where("word_ID = ?", param).Count(&count).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountFavoriteWordsU recieves a user ID, returns the number of words that user has marked
// it as favorite.
func CountFavoriteWordsU(c *gin.Context) {
	var count int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("userID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	if err := models.DB.Model(&models.FavoriteWord{}).Where("user_ID = ?", param).
		Count(&count).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// DeleteFavoriteWords deletes a favoriteWords
func DeleteFavoriteWords(c *gin.Context) {
	// Get model if exist
	var favoriteWords models.FavoriteWord
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("userID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("wordID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("user_ID = ? AND word_ID = ?", param1, param2).
		First(&favoriteWords).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&favoriteWords)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
