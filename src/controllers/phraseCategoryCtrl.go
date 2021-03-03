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
		phraseCategories[i].NameDe = security.RemoveBackticks(phraseCategories[i].NameDe)
		phraseCategories[i].NameEs = security.RemoveBackticks(phraseCategories[i].NameEs)
		phraseCategories[i].NameEn = security.RemoveBackticks(phraseCategories[i].NameEn)
		phraseCategories[i].NameFr = security.RemoveBackticks(phraseCategories[i].NameFr)
		phraseCategories[i].NameIt = security.RemoveBackticks(phraseCategories[i].NameIt)
		phraseCategories[i].NamePt = security.RemoveBackticks(phraseCategories[i].NamePt)
	}

	c.JSON(http.StatusOK, gin.H{"data": phraseCategories})
}

// CreatePhraseCategory creates a new phraseCategory.
func CreatePhraseCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.PhraseCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	phraseCategory := models.PhraseCategory{
		NameDe:   security.SecureString(input.NameDe),
		NameEs:   security.SecureString(input.NameEs),
		NameEn:   security.SecureString(input.NameEn),
		NameFr:   security.SecureString(input.NameFr),
		NameIt:   security.SecureString(input.NameIt),
		NamePt:   security.SecureString(input.NamePt),
		Modified: t}
	models.DB.Create(&phraseCategory)

	phraseCategory.NameDe = security.RemoveBackticks(phraseCategory.NameDe)
	phraseCategory.NameEs = security.RemoveBackticks(phraseCategory.NameEs)
	phraseCategory.NameEn = security.RemoveBackticks(phraseCategory.NameEn)
	phraseCategory.NameFr = security.RemoveBackticks(phraseCategory.NameFr)
	phraseCategory.NameIt = security.RemoveBackticks(phraseCategory.NameIt)
	phraseCategory.NamePt = security.RemoveBackticks(phraseCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// FindPhraseCategory recieves an id, and returns an specific phraseCategory with that id.
func FindPhraseCategory(c *gin.Context) {
	var phraseCategory models.PhraseCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}

	phraseCategory.NameDe = security.RemoveBackticks(phraseCategory.NameDe)
	phraseCategory.NameEs = security.RemoveBackticks(phraseCategory.NameEs)
	phraseCategory.NameEn = security.RemoveBackticks(phraseCategory.NameEn)
	phraseCategory.NameFr = security.RemoveBackticks(phraseCategory.NameFr)
	phraseCategory.NameIt = security.RemoveBackticks(phraseCategory.NameIt)
	phraseCategory.NamePt = security.RemoveBackticks(phraseCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// PatchPhraseCategory updates a phraseCategory
func PatchPhraseCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var phraseCategory models.PhraseCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhraseCategory not found!"})
		return
	}

	var input models.PhraseCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&phraseCategory).Updates(
		models.PhraseCategory{
			ID:       phraseCategory.ID,
			NameDe:   security.SecureString(input.NameDe),
			NameEs:   security.SecureString(input.NameEs),
			NameEn:   security.SecureString(input.NameEn),
			NameFr:   security.SecureString(input.NameFr),
			NameIt:   security.SecureString(input.NameIt),
			NamePt:   security.SecureString(input.NamePt),
			Modified: t,
		})

	phraseCategory.NameDe = security.RemoveBackticks(phraseCategory.NameDe)
	phraseCategory.NameEs = security.RemoveBackticks(phraseCategory.NameEs)
	phraseCategory.NameEn = security.RemoveBackticks(phraseCategory.NameEn)
	phraseCategory.NameFr = security.RemoveBackticks(phraseCategory.NameFr)
	phraseCategory.NameIt = security.RemoveBackticks(phraseCategory.NameIt)
	phraseCategory.NamePt = security.RemoveBackticks(phraseCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": phraseCategory})
}

// DeletePhraseCategory deletes a phraseCategory
func DeletePhraseCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var phraseCategory models.PhraseCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phraseCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&phraseCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
