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
		wordCategories[i].NameDe = security.RemoveBackticks(wordCategories[i].NameDe)
		wordCategories[i].NameEs = security.RemoveBackticks(wordCategories[i].NameEs)
		wordCategories[i].NameEn = security.RemoveBackticks(wordCategories[i].NameEn)
		wordCategories[i].NameFr = security.RemoveBackticks(wordCategories[i].NameFr)
		wordCategories[i].NameIt = security.RemoveBackticks(wordCategories[i].NameIt)
		wordCategories[i].NamePt = security.RemoveBackticks(wordCategories[i].NamePt)
	}

	c.JSON(http.StatusOK, gin.H{"data": wordCategories})
}

// CreateWordCategory creates a new wordCategory.
func CreateWordCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.WordCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	wordCategory := models.WordCategory{
		NameDe:   security.SecureString(input.NameDe),
		NameEs:   security.SecureString(input.NameEs),
		NameEn:   security.SecureString(input.NameEn),
		NameFr:   security.SecureString(input.NameFr),
		NameIt:   security.SecureString(input.NameIt),
		NamePt:   security.SecureString(input.NamePt),
		Modified: t}
	models.DB.Create(&wordCategory)

	wordCategory.NameDe = security.RemoveBackticks(wordCategory.NameDe)
	wordCategory.NameEs = security.RemoveBackticks(wordCategory.NameEs)
	wordCategory.NameEn = security.RemoveBackticks(wordCategory.NameEn)
	wordCategory.NameFr = security.RemoveBackticks(wordCategory.NameFr)
	wordCategory.NameIt = security.RemoveBackticks(wordCategory.NameIt)
	wordCategory.NamePt = security.RemoveBackticks(wordCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// FindWordCategory recieves an id, and returns an specific wordCategory with that id.
func FindWordCategory(c *gin.Context) {
	var wordCategory models.WordCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordCategory not found!"})
		return
	}

	wordCategory.NameDe = security.RemoveBackticks(wordCategory.NameDe)
	wordCategory.NameEs = security.RemoveBackticks(wordCategory.NameEs)
	wordCategory.NameEn = security.RemoveBackticks(wordCategory.NameEn)
	wordCategory.NameFr = security.RemoveBackticks(wordCategory.NameFr)
	wordCategory.NameIt = security.RemoveBackticks(wordCategory.NameIt)
	wordCategory.NamePt = security.RemoveBackticks(wordCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// PatchWordCategory updates a wordCategory
func PatchWordCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var wordCategory models.WordCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WordCategory not found!"})
		return
	}
	var input models.WordCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&wordCategory).Updates(
		models.WordCategory{
			ID:       wordCategory.ID,
			NameDe:   security.SecureString(input.NameDe),
			NameEs:   security.SecureString(input.NameEs),
			NameEn:   security.SecureString(input.NameEn),
			NameFr:   security.SecureString(input.NameFr),
			NameIt:   security.SecureString(input.NameIt),
			NamePt:   security.SecureString(input.NamePt),
			Modified: t,
		})

	wordCategory.NameDe = security.RemoveBackticks(wordCategory.NameDe)
	wordCategory.NameEs = security.RemoveBackticks(wordCategory.NameEs)
	wordCategory.NameEn = security.RemoveBackticks(wordCategory.NameEn)
	wordCategory.NameFr = security.RemoveBackticks(wordCategory.NameFr)
	wordCategory.NameIt = security.RemoveBackticks(wordCategory.NameIt)
	wordCategory.NamePt = security.RemoveBackticks(wordCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": wordCategory})
}

// DeleteWordCategory deletes a wordCategory
func DeleteWordCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var wordCategory models.WordCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&wordCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&wordCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
