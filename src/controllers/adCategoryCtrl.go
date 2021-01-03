package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAdCategories retrieves all the adCategories from the DB.
func GetAdCategories(c *gin.Context) {
	var adCategories []models.AdCategory
	models.DB.Find(&adCategories)

	for i := range adCategories {
		adCategories[i].Name = security.RemoveBackticks(adCategories[i].Name)
	}

	c.JSON(http.StatusOK, gin.H{"data": adCategories})
}

// CreateAdCategory creates a new adCategory.
func CreateAdCategory(c *gin.Context) {
	var input models.CreateAdCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	adCategory := models.AdCategory{
		Name:     security.SecureString(input.Name),
		Cost:     uint(input.Cost),
		Modified: t}
	models.DB.Create(&adCategory)

	adCategory.Name = security.RemoveBackticks(adCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// FindAdCategory recieves an id, and returns an specific adCategory with that id.
func FindAdCategory(c *gin.Context) {
	var adCategory models.AdCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AdCategory not found!"})
		return
	}

	adCategory.Name = security.RemoveBackticks(adCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// PatchAdCategory updates a adCategory
func PatchAdCategory(c *gin.Context) {

	// Get model if exist
	var adCategory models.AdCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AdCategory not found!"})
		return
	}

	var input models.CreateAdCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&adCategory).Updates(
		models.AdCategory{
			ID:       adCategory.ID,
			Name:     security.SecureString(input.Name),
			Cost:     uint(input.Cost),
			Modified: t,
		})

	adCategory.Name = security.RemoveBackticks(adCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// DeleteAdCategory deletes a adCategory
func DeleteAdCategory(c *gin.Context) {
	// Get model if exist
	var adCategory models.AdCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&adCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
