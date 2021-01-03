package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPhraseCategories retrieves all the phraseCategories from the DB.
func GetPhraseCategories(c *gin.Context) {
	var phraseCategories []models.PhraseCategory
	models.DB.Find(&phraseCategories)

	for i := range phraseCategories {
		phraseCategories[i].Name = security.RemoveBackticks(phraseCategories[i].Name)
	}

	c.JSON(http.StatusOK, gin.H{"data": phraseCategories})
}

// CreatePhraseCategory creates a new phraseCategory.
func CreatePhraseCategory(c *gin.Context) {
	var input models.CreatePhraseCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	phraseCategory := models.PhraseCategory{
		Name:     security.SecureString(input.Name),
		Modified: t}
	models.DB.Create(&phraseCategory)

	phraseCategory.Name = security.RemoveBackticks(phraseCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// FindPhraseCategory recieves an id, and returns an specific phraseCategory with that id.
func FindPhraseCategory(c *gin.Context) {
	var phraseCategory models.PhraseCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}

	phraseCategory.Name = security.RemoveBackticks(phraseCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// PatchPhraseCategory updates a phraseCategory
func PatchPhraseCategory(c *gin.Context) {

	// Get model if exist
	var phraseCategory models.PhraseCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}

	var input models.CreatePhraseCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&phraseCategory).Updates(
		models.PhraseCategory{
			ID:       phraseCategory.ID,
			Name:     security.SecureString(input.Name),
			Modified: t,
		})

	phraseCategory.Name = security.RemoveBackticks(phraseCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// DeletePhraseCategory deletes a phraseCategory
func DeletePhraseCategory(c *gin.Context) {
	// Get model if exist
	var phraseCategory models.PhraseCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&phraseCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
