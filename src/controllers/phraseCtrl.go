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
		phrases[i].TextDe = security.RemoveBackticks(phrases[i].TextDe)
		phrases[i].TextEs = security.RemoveBackticks(phrases[i].TextEs)
		phrases[i].TextEn = security.RemoveBackticks(phrases[i].TextEn)
		phrases[i].TextFr = security.RemoveBackticks(phrases[i].TextFr)
		phrases[i].TextIt = security.RemoveBackticks(phrases[i].TextIt)
		phrases[i].TextPt = security.RemoveBackticks(phrases[i].TextPt)
		phrases[i].ContextDe = security.RemoveBackticks(phrases[i].ContextDe)
		phrases[i].ContextEs = security.RemoveBackticks(phrases[i].ContextEs)
		phrases[i].ContextEn = security.RemoveBackticks(phrases[i].ContextEn)
		phrases[i].ContextFr = security.RemoveBackticks(phrases[i].ContextFr)
		phrases[i].ContextIt = security.RemoveBackticks(phrases[i].ContextIt)
		phrases[i].ContextPt = security.RemoveBackticks(phrases[i].ContextPt)
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
		TextDe:           security.SecureString(input.TextDe),
		TextEs:           security.SecureString(input.TextEs),
		TextEn:           security.SecureString(input.TextEn),
		TextFr:           security.SecureString(input.TextFr),
		TextIt:           security.SecureString(input.TextIt),
		TextPt:           security.SecureString(input.TextPt),
		ContextDe:        security.SecureString(input.ContextDe),
		ContextEs:        security.SecureString(input.ContextEs),
		ContextEn:        security.SecureString(input.ContextEn),
		ContextFr:        security.SecureString(input.ContextFr),
		ContextIt:        security.SecureString(input.ContextIt),
		ContextPt:        security.SecureString(input.ContextPt),
		Modified:         t,
	}
	models.DB.Create(&phrase)

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.TextDe = security.RemoveBackticks(phrase.TextDe)
	phrase.TextEs = security.RemoveBackticks(phrase.TextEs)
	phrase.TextEn = security.RemoveBackticks(phrase.TextEn)
	phrase.TextFr = security.RemoveBackticks(phrase.TextFr)
	phrase.TextIt = security.RemoveBackticks(phrase.TextIt)
	phrase.TextPt = security.RemoveBackticks(phrase.TextPt)
	phrase.ContextDe = security.RemoveBackticks(phrase.ContextDe)
	phrase.ContextEs = security.RemoveBackticks(phrase.ContextEs)
	phrase.ContextEn = security.RemoveBackticks(phrase.ContextEn)
	phrase.ContextFr = security.RemoveBackticks(phrase.ContextFr)
	phrase.ContextIt = security.RemoveBackticks(phrase.ContextIt)
	phrase.ContextPt = security.RemoveBackticks(phrase.ContextPt)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// FindPhrase recieves an id, and returns an specific phrase with that id.
func FindPhrase(c *gin.Context) {
	var phrase models.Phrase
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrase not found!"})
		return
	}

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.TextDe = security.RemoveBackticks(phrase.TextDe)
	phrase.TextEs = security.RemoveBackticks(phrase.TextEs)
	phrase.TextEn = security.RemoveBackticks(phrase.TextEn)
	phrase.TextFr = security.RemoveBackticks(phrase.TextFr)
	phrase.TextIt = security.RemoveBackticks(phrase.TextIt)
	phrase.TextPt = security.RemoveBackticks(phrase.TextPt)
	phrase.ContextDe = security.RemoveBackticks(phrase.ContextDe)
	phrase.ContextEs = security.RemoveBackticks(phrase.ContextEs)
	phrase.ContextEn = security.RemoveBackticks(phrase.ContextEn)
	phrase.ContextFr = security.RemoveBackticks(phrase.ContextFr)
	phrase.ContextIt = security.RemoveBackticks(phrase.ContextIt)
	phrase.ContextPt = security.RemoveBackticks(phrase.ContextPt)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// PatchPhrase updates a phrase
func PatchPhrase(c *gin.Context) {

	// Get model if exist
	var phrase models.Phrase
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phrase not found!"})
		return
	}

	var input models.UpdatePhraseInput

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
			TextDe:           security.SecureString(input.TextDe),
			TextEs:           security.SecureString(input.TextEs),
			TextEn:           security.SecureString(input.TextEn),
			TextFr:           security.SecureString(input.TextFr),
			TextIt:           security.SecureString(input.TextIt),
			TextPt:           security.SecureString(input.TextPt),
			ContextDe:        security.SecureString(input.ContextDe),
			ContextEs:        security.SecureString(input.ContextEs),
			ContextEn:        security.SecureString(input.ContextEn),
			ContextFr:        security.SecureString(input.ContextFr),
			ContextIt:        security.SecureString(input.ContextIt),
			ContextPt:        security.SecureString(input.ContextPt),
			Modified:         t,
		})

	phrase.Text = security.RemoveBackticks(phrase.Text)
	phrase.TextDe = security.RemoveBackticks(phrase.TextDe)
	phrase.TextEs = security.RemoveBackticks(phrase.TextEs)
	phrase.TextEn = security.RemoveBackticks(phrase.TextEn)
	phrase.TextFr = security.RemoveBackticks(phrase.TextFr)
	phrase.TextIt = security.RemoveBackticks(phrase.TextIt)
	phrase.TextPt = security.RemoveBackticks(phrase.TextPt)
	phrase.ContextDe = security.RemoveBackticks(phrase.ContextDe)
	phrase.ContextEs = security.RemoveBackticks(phrase.ContextEs)
	phrase.ContextEn = security.RemoveBackticks(phrase.ContextEn)
	phrase.ContextFr = security.RemoveBackticks(phrase.ContextFr)
	phrase.ContextIt = security.RemoveBackticks(phrase.ContextIt)
	phrase.ContextPt = security.RemoveBackticks(phrase.ContextPt)

	c.JSON(http.StatusOK, gin.H{"data": phrase})
}

// DeletePhrase deletes a phrase
func DeletePhrase(c *gin.Context) {
	// Get model if exist
	var phrase models.Phrase
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&phrase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&phrase)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
