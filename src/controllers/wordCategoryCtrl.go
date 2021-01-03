package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWordCategories retrieves all the wordCategories from the DB.
func GetWordCategories(c *gin.Context) {
	var wordCategories []models.WordCategory
	models.DB.Find(&wordCategories)

	for i := range wordCategories {
		wordCategories[i].Name = security.RemoveBackticks(wordCategories[i].Name)
	}

	c.JSON(http.StatusOK, gin.H{"data": wordCategories})
}

// CreateWordCategory creates a new wordCategory.
func CreateWordCategory(c *gin.Context) {
	var input models.CreateWordCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	wordCategory := models.WordCategory{
		Name:     security.SecureString(input.Name),
		Modified: t}
	models.DB.Create(&wordCategory)

	wordCategory.Name = security.RemoveBackticks(wordCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// FindWordCategory recieves an id, and returns an specific wordCategory with that id.
func FindWordCategory(c *gin.Context) {
	var wordCategory models.WordCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordCategory not found!"})
		return
	}

	wordCategory.Name = security.RemoveBackticks(wordCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// PatchWordCategory updates a wordCategory
func PatchWordCategory(c *gin.Context) {
	// Get model if exist
	var wordCategory models.WordCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordCategory not found!"})
		return
	}
	var input models.CreateWordCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&wordCategory).Updates(
		models.WordCategory{
			ID:       wordCategory.ID,
			Name:     security.SecureString(input.Name),
			Modified: t,
		})

	wordCategory.Name = security.RemoveBackticks(wordCategory.Name)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// DeleteWordCategory deletes a wordCategory
func DeleteWordCategory(c *gin.Context) {
	// Get model if exist
	var wordCategory models.WordCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&wordCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
