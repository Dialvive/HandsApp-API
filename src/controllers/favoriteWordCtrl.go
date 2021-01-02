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

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	favoriteWord := models.FavoriteWord{
		WordID:   input.WordID,
		UserID:   input.UserID,
		Modified: t,
	}
	models.DB.Create(&favoriteWord)

	c.JSON(http.StatusOK, gin.H{"data": favoriteWord})
}

// FindFavoriteWords recieves an user id, and returns all of its favorite words.
func FindFavoriteWords(c *gin.Context) {
	var favorites []models.FavoriteWord
	if err := models.DB.Where("user_ID = ?", c.Param("id")).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Words not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountFavoriteWordsP recieves a word ID, returns the number of users that user has marked it as favorite.
func CountFavoriteWordsP(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoriteWord{}).Where("word_ID = ?", c.Param("wordId")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountFavoriteWordsU recieves a user ID, returns the number of words that user has marked it as favorite.
func CountFavoriteWordsU(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.FavoriteWord{}).Where("user_ID = ?", c.Param("userId")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// PutFavoriteWord updates a favoriteWord
func PutFavoriteWord(c *gin.Context) {

	// Get model if exist
	var favoriteWord models.FavoriteWord
	if err := models.DB.Where("user_ID = ? AND word_ID = ?", c.Param("userId"), c.Param("wordId")).First(&favoriteWord).Error; err != nil {
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
	if err := models.DB.Where("user_ID = ? AND word_ID = ?", c.Param("userId"), c.Param("word_Id")).First(&favoriteWord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&favoriteWord)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
