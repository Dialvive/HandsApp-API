package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPhrases retrieves all the phrases from the DB.
func GetPhrases(c *gin.Context) {
	var phrases []models.Phrase
	models.DB.Find(&phrases)

	for i := range phrases {
		phrases[i].Text = security.RemoveBackticks(phrases[i].Text)
		phrases[i].Context = security.RemoveBackticks(phrases[i].Context)
	}

	c.JSON(http.StatusOK, gin.H{"data": phrases})
}

// CreatePhrase creates a new phrase.
func CreatePhrase(c *gin.Context) {
	var input models.CreatePhraseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	phrase := models.Phrase{
		LocaleID:         uint(input.LocaleID),
		PhraseCategoryID: uint(input.PhraseCategoryID),
		Text:             security.SecureString(input.Text),
		Context:          security.SecureString(input.Context),
		Modified:         t,
	}
	models.DB.Create(&phrase)

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.Context = security.RemoveBackticks(phrase.Context)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// FindPhrase recieves an id, and returns an specific phrase with that id.
func FindPhrase(c *gin.Context) {
	var phrase models.Phrase

	if err := models.DB.Where("id = ?", c.Param("id")).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrase not found!"})
		return
	}

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.Context = security.RemoveBackticks(phrase.Context)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// PatchPhrase updates a phrase
func PatchPhrase(c *gin.Context) {

	// Get model if exist
	var phrase models.Phrase

	if err := models.DB.Where("id = ?", c.Param("id")).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrase not found!"})
		return
	}

	var input models.CreatePhraseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&phrase).Updates(
		models.Phrase{
			ID:               phrase.ID,
			LocaleID:         uint(input.LocaleID),
			PhraseCategoryID: uint(input.PhraseCategoryID),
			Text:             security.SecureString(input.Text),
			Context:          security.SecureString(input.Context),
			Modified:         t,
		})

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.Context = security.RemoveBackticks(phrase.Context)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// DeletePhrase deletes a phrase
func DeletePhrase(c *gin.Context) {
	// Get model if exist
	var phrase models.Phrase
	if err := models.DB.Where("id = ?", c.Param("id")).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&phrase)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
