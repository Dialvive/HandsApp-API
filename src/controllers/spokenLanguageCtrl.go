package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSpokenLanguages retrieves all the spokenLanguages from the DB.
func GetSpokenLanguages(c *gin.Context) {
	var spokenLanguages []models.SpokenLanguage
	models.DB.Find(&spokenLanguages)

	for i := range spokenLanguages {
		spokenLanguages[i].NameDe = security.RemoveBackticks(spokenLanguages[i].NameDe)
		spokenLanguages[i].NameEs = security.RemoveBackticks(spokenLanguages[i].NameEs)
		spokenLanguages[i].NameEn = security.RemoveBackticks(spokenLanguages[i].NameEn)
		spokenLanguages[i].NameFr = security.RemoveBackticks(spokenLanguages[i].NameFr)
		spokenLanguages[i].NameIt = security.RemoveBackticks(spokenLanguages[i].NameIt)
		spokenLanguages[i].NamePt = security.RemoveBackticks(spokenLanguages[i].NamePt)
		spokenLanguages[i].Abbreviation = security.RemoveBackticks(spokenLanguages[i].Abbreviation)
	}

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
		NameDe:       security.SecureString(input.NameDe),
		NameEs:       security.SecureString(input.NameEs),
		NameEn:       security.SecureString(input.NameEn),
		NameFr:       security.SecureString(input.NameFr),
		NameIt:       security.SecureString(input.NameIt),
		NamePt:       security.SecureString(input.NamePt),
		Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 2)),
		Modified:     t}
	models.DB.Create(&spokenLanguage)

	spokenLanguage.NameDe = security.RemoveBackticks(spokenLanguage.NameDe)
	spokenLanguage.NameEs = security.RemoveBackticks(spokenLanguage.NameEs)
	spokenLanguage.NameEn = security.RemoveBackticks(spokenLanguage.NameEn)
	spokenLanguage.NameFr = security.RemoveBackticks(spokenLanguage.NameFr)
	spokenLanguage.NameIt = security.RemoveBackticks(spokenLanguage.NameIt)
	spokenLanguage.NamePt = security.RemoveBackticks(spokenLanguage.NamePt)
	spokenLanguage.Abbreviation = security.RemoveBackticks(spokenLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// FindSpokenLanguage recieves an id, and returns an specific spokenLanguage with that id.
func FindSpokenLanguage(c *gin.Context) {
	var spokenLanguage models.SpokenLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SpokenLanguage not found!"})
		return
	}
	spokenLanguage.NameDe = security.RemoveBackticks(spokenLanguage.NameDe)
	spokenLanguage.NameEs = security.RemoveBackticks(spokenLanguage.NameEs)
	spokenLanguage.NameEn = security.RemoveBackticks(spokenLanguage.NameEn)
	spokenLanguage.NameFr = security.RemoveBackticks(spokenLanguage.NameFr)
	spokenLanguage.NameIt = security.RemoveBackticks(spokenLanguage.NameIt)
	spokenLanguage.NamePt = security.RemoveBackticks(spokenLanguage.NamePt)
	spokenLanguage.Abbreviation = security.RemoveBackticks(spokenLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// PatchSpokenLanguage updates a spokenLanguage
func PatchSpokenLanguage(c *gin.Context) {

	// Get model if exist
	var spokenLanguage models.SpokenLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SpokenLanguage not found!"})
		return
	}

	var input models.UpdateSpokenLanguageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&spokenLanguage).Updates(
		models.SpokenLanguage{
			ID:           spokenLanguage.ID,
			NameDe:       security.SecureString(input.NameDe),
			NameEs:       security.SecureString(input.NameEs),
			NameEn:       security.SecureString(input.NameEn),
			NameFr:       security.SecureString(input.NameFr),
			NameIt:       security.SecureString(input.NameIt),
			NamePt:       security.SecureString(input.NamePt),
			Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 2)),
			Modified:     t,
		})

	spokenLanguage.NameDe = security.RemoveBackticks(spokenLanguage.NameDe)
	spokenLanguage.NameEs = security.RemoveBackticks(spokenLanguage.NameEs)
	spokenLanguage.NameEn = security.RemoveBackticks(spokenLanguage.NameEn)
	spokenLanguage.NameFr = security.RemoveBackticks(spokenLanguage.NameFr)
	spokenLanguage.NameIt = security.RemoveBackticks(spokenLanguage.NameIt)
	spokenLanguage.NamePt = security.RemoveBackticks(spokenLanguage.NamePt)
	spokenLanguage.Abbreviation = security.RemoveBackticks(spokenLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": spokenLanguage})
}

// DeleteSpokenLanguage deletes a spokenLanguage
func DeleteSpokenLanguage(c *gin.Context) {
	// Get model if exist
	var spokenLanguage models.SpokenLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&spokenLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&spokenLanguage)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
