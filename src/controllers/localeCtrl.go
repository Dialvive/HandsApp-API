package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetLocales retrieves all the locales from the DB.
func GetLocales(c *gin.Context) {
	var locales []models.Locale
	models.DB.Find(&locales)

	c.JSON(http.StatusOK, gin.H{"data": locales})
}

// CreateLocale creates a new locale.
func CreateLocale(c *gin.Context) {
	var input models.CreateLocaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	locale := models.Locale{
		CountryID:        uint(input.CountryID),
		SpokenLanguageID: uint(input.SpokenLanguageID),
		SignLanguageID:   uint(input.SignLanguageID),
		Modified:         t,
	}
	models.DB.Create(&locale)

	c.JSON(http.StatusOK, gin.H{"data": locale})
}

// FindLocale recieves an id, and returns an specific locale with that id.
func FindLocale(c *gin.Context) {
	var locale models.Locale

	if err := models.DB.Where("id = ?", c.Param("id")).First(&locale).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Locale not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": locale})
}

// PatchLocale updates a locale
func PatchLocale(c *gin.Context) {

	// Get model if exist
	var locale models.Locale

	if err := models.DB.Where("id = ?", c.Param("id")).First(&locale).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Locale not found!"})
		return
	}

	var input models.CreateLocaleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&locale).Updates(
		models.Locale{
			ID:               locale.ID,
			CountryID:        uint(input.CountryID),
			SpokenLanguageID: uint(input.SpokenLanguageID),
			SignLanguageID:   uint(input.SignLanguageID),
			Modified:         t,
		})

	c.JSON(http.StatusOK, gin.H{"data": locale})
}

// DeleteLocale deletes a locale
func DeleteLocale(c *gin.Context) {
	// Get model if exist
	var locale models.Locale
	if err := models.DB.Where("id = ?", c.Param("id")).First(&locale).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&locale)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
