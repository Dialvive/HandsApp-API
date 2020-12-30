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
		WordID:   input.WordID,
		RegionID: input.RegionID,
		Modified: t,
	}
	models.DB.Create(&wordByRegion)

	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// FindWordByRegion recieves an id, and returns an specific wordByRegion with that id.
func FindWordByRegion(c *gin.Context) {
	var wordByRegion models.WordByRegion

	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordByRegion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordByRegion not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// PutWordByRegion updates a wordByRegion
func PutWordByRegion(c *gin.Context) {

	// Get model if exist
	var wordByRegion models.WordByRegion

	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordByRegion).Error; err != nil {
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
			WordID:   input.WordID,
			RegionID: input.RegionID,
			Modified: t,
		})

	c.JSON(http.StatusOK, gin.H{"data": wordByRegion})
}

// DeleteWordByRegion deletes a wordByRegion
func DeleteWordByRegion(c *gin.Context) {
	// Get model if exist
	var wordByRegion models.WordByRegion
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordByRegion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&wordByRegion)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
