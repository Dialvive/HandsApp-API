package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSignLanguages retrieves all the signLanguages from the DB.
func GetSignLanguages(c *gin.Context) {
	var signLanguages []models.SignLanguage
	models.DB.Find(&signLanguages)

	for i := range signLanguages {
		signLanguages[i].Name = security.RemoveBackticks(signLanguages[i].Name)
		signLanguages[i].Abbreviation = security.RemoveBackticks(signLanguages[i].Abbreviation)
	}

	c.JSON(http.StatusOK, gin.H{"data": signLanguages})
}

// CreateSignLanguage creates a new signLanguage.
func CreateSignLanguage(c *gin.Context) {
	var input models.CreateSignLanguageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	signLanguage := models.SignLanguage{
		Name:         security.SecureString(input.Name),
		Abbreviation: security.SecureString(input.Abbreviation),
		Modified:     t}
	models.DB.Create(&signLanguage)

	signLanguage.Name = security.RemoveBackticks(signLanguage.Name)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// FindSignLanguage recieves an id, and returns an specific signLanguage with that id.
func FindSignLanguage(c *gin.Context) {
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignLanguage not found!"})
		return
	}

	signLanguage.Name = security.RemoveBackticks(signLanguage.Name)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// PatchSignLanguage updates a signLanguage
func PatchSignLanguage(c *gin.Context) {

	// Get model if exist
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignLanguage not found!"})
		return
	}

	var input models.CreateSignLanguageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&signLanguage).Updates(
		models.SignLanguage{
			ID:           signLanguage.ID,
			Name:         security.SecureString(input.Name),
			Abbreviation: security.SecureString(input.Abbreviation),
			Modified:     t,
		})

	signLanguage.Name = security.RemoveBackticks(signLanguage.Name)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// DeleteSignLanguage deletes a signLanguage
func DeleteSignLanguage(c *gin.Context) {
	// Get model if exist
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&signLanguage)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
