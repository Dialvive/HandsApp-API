package controllers

import (
	"API/models"
	"API/security"
	"net/http"

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
	var input models.WordByRegionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wordByRegion := models.WordByRegion{
		WordID:   uint(input.WordID),
		RegionID: uint(input.RegionID),
	}
	models.DB.Create(&wordByRegion)

	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// FindWordsOfRegion recieves an user id, and returns all of its favorite words.
func FindWordsOfRegion(c *gin.Context) {
	var words []models.WordByRegion
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("regionID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("region_ID = ?", param).Find(&words).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Words not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": words})
}

// CountWordsOfRegion recieves a word ID, returns the number of users that user has marked
// it as favorite.
func CountWordsOfRegion(c *gin.Context) {
	var count int64
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("regionID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Model(&models.WordByRegion{}).Where("region_ID = ?", param).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// DeleteWordByRegion deletes a wordByRegion
func DeleteWordByRegion(c *gin.Context) {
	// Get model if exist
	var wordByRegion models.WordByRegion
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("regionID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if param2, err = security.SecureUint(c.Param("wordID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("region_ID = ? AND word_ID = ?", param1, param2).
		First(&wordByRegion).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&wordByRegion)
	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}
