package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRegions retrieves all the regions from the DB.
func GetRegions(c *gin.Context) {
	var regions []models.Region
	models.DB.Find(&regions)

	for i := range regions {
		regions[i].Name = security.RemoveBackticks(regions[i].Name)
	}

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
		Name:      security.SecureString(input.Name),
		CountryID: uint(input.CountryID),
		Modified:  t,
	}
	models.DB.Create(&region)

	region.Name = security.RemoveBackticks(region.Name)

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// FindRegion recieves an id, and returns an specific region with that id.
func FindRegion(c *gin.Context) {
	var region models.Region
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region not found!"})
		return
	}

	region.Name = security.RemoveBackticks(region.Name)

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// PatchRegion updates a region
func PatchRegion(c *gin.Context) {

	// Get model if exist
	var region models.Region
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region not found!"})
		return
	}

	var input models.UpdateRegionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&region).Updates(
		models.Region{
			ID:        region.ID,
			Name:      security.SecureString(input.Name),
			CountryID: uint(input.CountryID),
			Modified:  t,
		})

	region.Name = security.RemoveBackticks(region.Name)

	c.JSON(http.StatusOK, gin.H{"data": region})
}

// DeleteRegion deletes a region
func DeleteRegion(c *gin.Context) {
	// Get model if exist
	var region models.Region
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&region)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
