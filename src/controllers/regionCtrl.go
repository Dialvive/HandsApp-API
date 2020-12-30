package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRegions retrieves all the regions from the DB.
func GetRegions(c *gin.Context) {
	var regions []models.Region
	models.DB.Find(&regions)

	c.JSON(http.StatusOK, gin.H{"data": regions})
}

// CreateRegion creates a new region.
func CreateRegion(c *gin.Context) {
	var input models.CreateRegionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	region := models.Region{
		Name:      input.Name,
		CountryID: uint(input.CountryID),
		Modified:  t,
	}
	models.DB.Create(&region)

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// FindRegion recieves an id, and returns an specific region with that id.
func FindRegion(c *gin.Context) {
	var region models.Region

	if err := models.DB.Where("id = ?", c.Param("id")).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// PatchRegion updates a region
func PatchRegion(c *gin.Context) {

	// Get model if exist
	var region models.Region

	if err := models.DB.Where("id = ?", c.Param("id")).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region not found!"})
		return
	}

	var input models.CreateRegionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&region).Updates(
		models.Region{
			ID:        region.ID,
			Name:      input.Name,
			CountryID: uint(input.CountryID),
			Modified:  t,
		})

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// DeleteRegion deletes a region
func DeleteRegion(c *gin.Context) {
	// Get model if exist
	var region models.Region
	if err := models.DB.Where("id = ?", c.Param("id")).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&region)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
