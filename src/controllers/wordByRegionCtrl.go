package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWordsByRegions retrieves all the wordsByRegions from the DB.
func GetWordsByRegions(c *gin.Context) {
	var wordsByRegions []models.WordByRegion
	models.DB.Find(&wordsByRegions)

	c.JSON(http.StatusOK, gin.H{"data": wordsByRegions})
}

// CreateWordByRegion creates a new wordByRegion.
func CreateWordByRegion(c *gin.Context) {
	var input models.CreateWordByRegionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	wordByRegion := models.WordByRegion{
		WordID:   uint(input.WordID),
		RegionID: uint(input.RegionID),
		Modified: t,
	}
	models.DB.Create(&wordByRegion)

	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// FindWordsOfRegion recieves an user id, and returns all of its favorite words.
func FindWordsOfRegion(c *gin.Context) {
	var favorites []models.WordByRegion
	if err := models.DB.Where("user_ID = ?", c.Param("id")).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Words not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}

// CountWordsOfRegion recieves a word ID, returns the number of users that user has marked it as favorite.
func CountWordsOfRegion(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.WordByRegion{}).Where("word_ID = ?", c.Param("wordID")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// CountWordsByRegionU recieves a user ID, returns the number of words that user has marked it as favorite.
func CountWordsByRegionU(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.WordByRegion{}).Where("region_ID = ?", c.Param("regionID")).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// PutWordByRegion updates a wordByRegion
func PutWordByRegion(c *gin.Context) {

	// Get model if exist
	var wordByRegion models.WordByRegion
	if err := models.DB.Where("region_ID = ? AND word_ID = ?", c.Param("regionID"), c.Param("word_ID")).First(&wordByRegion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordByRegion not found!"})
		return
	}

	var input models.CreateWordByRegionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&wordByRegion).Updates(
		models.WordByRegion{
			WordID:   uint(input.WordID),
			RegionID: uint(input.RegionID),
			Modified: t,
		})
	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// DeleteWordByRegion deletes a wordByRegion
func DeleteWordByRegion(c *gin.Context) {
	// Get model if exist
	var wordByRegion models.WordByRegion
	if err := models.DB.Where("region_ID = ? AND word_ID = ?", c.Param("regionID"), c.Param("word_ID")).First(&wordByRegion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&wordByRegion)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
