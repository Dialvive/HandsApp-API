package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAdvertisements retrieves all the advertisements from the DB.
func GetAdvertisements(c *gin.Context) {
	var advertisements []models.Advertisement
	models.DB.Find(&advertisements)

	c.JSON(http.StatusOK, gin.H{"data": advertisements})
}

// CreateAdvertisement creates a new advertisement.
func CreateAdvertisement(c *gin.Context) {
	var input models.CreateAdvertisementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	advertisement := models.Advertisement{
		UserID:       uint(input.UserID),
		RegionID:     uint(input.RegionID),
		AdCategoryID: uint(input.AdCategoryID),
		Title:        input.Title,
		Body:         input.Body,
		Media:        input.Media,
		Paid:         uint(input.Paid),
		Modified:     t,
	}
	models.DB.Create(&advertisement)

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// FindAdvertisement recieves an id, and returns an specific advertisement with that id.
func FindAdvertisement(c *gin.Context) {
	var advertisement models.Advertisement

	if err := models.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Advertisement not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// PatchAdvertisement updates a advertisement
func PatchAdvertisement(c *gin.Context) {

	// Get model if exist
	var advertisement models.Advertisement

	if err := models.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Advertisement not found!"})
		return
	}

	var input models.CreateAdvertisementInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&advertisement).Updates(
		models.Advertisement{
			ID:           advertisement.ID,
			UserID:       uint(input.UserID),
			RegionID:     uint(input.RegionID),
			AdCategoryID: uint(input.AdCategoryID),
			Title:        input.Title,
			Body:         input.Body,
			Media:        input.Media,
			Paid:         uint(input.Paid),
			Modified:     t,
		})

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// DeleteAdvertisement deletes a advertisement
func DeleteAdvertisement(c *gin.Context) {
	// Get model if exist
	var advertisement models.Advertisement
	if err := models.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&advertisement)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
