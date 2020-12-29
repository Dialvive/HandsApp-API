package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSpokenLanguages retrieves all the spokenLanguages from the DB.
func GetSpokenLanguages(c *gin.Context) {
	var spokenLanguages []models.SpokenLanguage
	models.DB.Find(&spokenLanguages)

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguages})
}

// CreateSpokenLanguage creates a new spokenLanguage.
func CreateSpokenLanguage(c *gin.Context) {
	var input models.CreateSpokenLanguageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	spokenLanguage := models.SpokenLanguage{
		Name:         input.Name,
		Abbreviation: input.Abbreviation,
		Creation:     t}
	models.DB.Create(&spokenLanguage)

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// FindSpokenLanguage recieves an id, and returns an specific spokenLanguage with that id.
func FindSpokenLanguage(c *gin.Context) {
	var spokenLanguage models.SpokenLanguage

	if err := models.DB.Where("id = ?", c.Param("id")).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SpokenLanguage not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// UpdateSpokenLanguage updates a spokenLanguage
func UpdateSpokenLanguage(c *gin.Context) {

	// Get model if exist
	var spokenLanguage models.SpokenLanguage

	if err := models.DB.Where("id = ?", c.Param("id")).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SpokenLanguage not found!"})
		return
	}

	var input models.CreateSpokenLanguageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&spokenLanguage).Updates(
		models.SpokenLanguage{
			ID:           spokenLanguage.ID,
			Name:         input.Name,
			Abbreviation: input.Abbreviation,
			Creation:     t,
		})

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// DeleteSpokenLanguage deletes a spokenLanguage
func DeleteSpokenLanguage(c *gin.Context) {
	// Get model if exist
	var spokenLanguage models.SpokenLanguage
	if err := models.DB.Where("id = ?", c.Param("id")).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&spokenLanguage)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
