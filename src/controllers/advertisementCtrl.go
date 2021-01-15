package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAdvertisements retrieves all the advertisements from the DB.
func GetAdvertisements(c *gin.Context) {
	var advertisements []models.Advertisement
	models.DB.Find(&advertisements)

	for i := range advertisements {
		advertisements[i].Body = security.RemoveBackticks(advertisements[i].Body)
		advertisements[i].Title = security.RemoveBackticks(advertisements[i].Title)
	}

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
		Title:        security.SecureString(security.TrimToLength(input.Title, 62)),
		Body:         security.SecureString(input.Body),
		Media:        bool(input.Media),
		Paid:         uint(input.Paid),
		Modified:     t,
	}
	models.DB.Create(&advertisement)

	advertisement.Body = security.RemoveBackticks(advertisement.Body)
	advertisement.Title = security.RemoveBackticks(advertisement.Title)

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// FindAdvertisement recieves an id, and returns an specific advertisement with that id.
func FindAdvertisement(c *gin.Context) {
	var advertisement models.Advertisement
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Advertisement not found!"})
		return
	}

	advertisement.Body = security.RemoveBackticks(advertisement.Body)
	advertisement.Title = security.RemoveBackticks(advertisement.Title)

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// PatchAdvertisement updates a advertisement
func PatchAdvertisement(c *gin.Context) {

	// Get model if exist
	var advertisement models.Advertisement
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Advertisement not found!"})
		return
	}

	var input models.UpdateAdvertisementInput

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
			Title:        security.SecureString(security.TrimToLength(input.Title, 62)),
			Body:         security.SecureString(input.Body),
			Media:        bool(input.Media),
			Paid:         uint(input.Paid),
			Modified:     t,
		})

	advertisement.Body = security.RemoveBackticks(advertisement.Body)
	advertisement.Title = security.RemoveBackticks(advertisement.Title)

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

// DeleteAdvertisement deletes a advertisement
func DeleteAdvertisement(c *gin.Context) {
	// Get model if exist
	var advertisement models.Advertisement
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&advertisement)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
