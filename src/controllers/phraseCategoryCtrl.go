package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPhraseCategories retrieves all the phraseCategories from the DB.
func GetPhraseCategories(c *gin.Context) {
	var phraseCategories []models.PhraseCategory
	models.DB.Find(&phraseCategories)

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
		Name:     input.Name,
		Creation: t}
	models.DB.Create(&phraseCategory)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// FindPhraseCategory recieves an id, and returns an specific phraseCategory with that id.
func FindPhraseCategory(c *gin.Context) {
	var phraseCategory models.PhraseCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// UpdatePhraseCategory updates a phraseCategory
func UpdatePhraseCategory(c *gin.Context) {
	println("UpdatePhraseCategory")
	// Get model if exist
	var phraseCategory models.PhraseCategory
	println("models.PhraseCategory")
	if err := models.DB.Where("id = ?", c.Param("id")).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}
	println("no error")
	var input models.CreatePhraseCategoryInput
	println("models.UpdatePhraseCategoryInput")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println("no error")
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&phraseCategory).Updates(
		models.PhraseCategory{
			ID:       phraseCategory.ID,
			Name:     input.Name,
			Creation: t,
		})
	println("Update")

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
