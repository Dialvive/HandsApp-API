package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWords retrieves all the words from the DB.
func GetWords(c *gin.Context) {
	var words []models.Word
	models.DB.Find(&words)

	for i := range words {
		words[i].TextDe = security.RemoveBackticks(words[i].TextDe)
		words[i].TextEs = security.RemoveBackticks(words[i].TextEs)
		words[i].TextEn = security.RemoveBackticks(words[i].TextEn)
		words[i].TextFr = security.RemoveBackticks(words[i].TextFr)
		words[i].TextIt = security.RemoveBackticks(words[i].TextIt)
		words[i].TextPt = security.RemoveBackticks(words[i].TextPt)
		words[i].ContextDe = security.RemoveBackticks(words[i].ContextDe)
		words[i].ContextEs = security.RemoveBackticks(words[i].ContextEs)
		words[i].ContextEn = security.RemoveBackticks(words[i].ContextEn)
		words[i].ContextFr = security.RemoveBackticks(words[i].ContextFr)
		words[i].ContextIt = security.RemoveBackticks(words[i].ContextIt)
		words[i].ContextPt = security.RemoveBackticks(words[i].ContextPt)
		words[i].DefinitionDe = security.RemoveBackticks(words[i].DefinitionDe)
		words[i].DefinitionEs = security.RemoveBackticks(words[i].DefinitionEs)
		words[i].DefinitionEn = security.RemoveBackticks(words[i].DefinitionEn)
		words[i].DefinitionFr = security.RemoveBackticks(words[i].DefinitionFr)
		words[i].DefinitionIt = security.RemoveBackticks(words[i].DefinitionIt)
		words[i].DefinitionPt = security.RemoveBackticks(words[i].DefinitionPt)
	}

	c.JSON(http.StatusOK, gin.H{"data": words})
}

// CreateWord creates a new word.
func CreateWord(c *gin.Context) {
	var input models.CreateWordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	word := models.Word{
		LocaleID:       uint(input.LocaleID),
		WordCategoryID: uint(input.WordCategoryID),
		TextDe:         security.SecureString(input.TextDe),
		TextEs:         security.SecureString(input.TextEs),
		TextEn:         security.SecureString(input.TextEn),
		TextFr:         security.SecureString(input.TextFr),
		TextIt:         security.SecureString(input.TextIt),
		TextPt:         security.SecureString(input.TextPt),
		ContextDe:      security.SecureString(input.ContextDe),
		ContextEs:      security.SecureString(input.ContextEs),
		ContextEn:      security.SecureString(input.ContextEn),
		ContextFr:      security.SecureString(input.ContextFr),
		ContextIt:      security.SecureString(input.ContextIt),
		ContextPt:      security.SecureString(input.ContextPt),
		DefinitionDe:   security.SecureString(input.DefinitionDe),
		DefinitionEs:   security.SecureString(input.DefinitionEs),
		DefinitionEn:   security.SecureString(input.DefinitionEn),
		DefinitionFr:   security.SecureString(input.DefinitionFr),
		DefinitionIt:   security.SecureString(input.DefinitionIt),
		DefinitionPt:   security.SecureString(input.DefinitionPt),
		Modified:       t,
	}
	models.DB.Create(&word)

	word.TextDe = security.RemoveBackticks(word.TextDe)
	word.TextEs = security.RemoveBackticks(word.TextEs)
	word.TextEn = security.RemoveBackticks(word.TextEn)
	word.TextFr = security.RemoveBackticks(word.TextFr)
	word.TextIt = security.RemoveBackticks(word.TextIt)
	word.TextPt = security.RemoveBackticks(word.TextPt)
	word.ContextDe = security.RemoveBackticks(word.ContextDe)
	word.ContextEs = security.RemoveBackticks(word.ContextEs)
	word.ContextEn = security.RemoveBackticks(word.ContextEn)
	word.ContextFr = security.RemoveBackticks(word.ContextFr)
	word.ContextIt = security.RemoveBackticks(word.ContextIt)
	word.ContextPt = security.RemoveBackticks(word.ContextPt)
	word.DefinitionDe = security.RemoveBackticks(word.DefinitionDe)
	word.DefinitionEs = security.RemoveBackticks(word.DefinitionEs)
	word.DefinitionEn = security.RemoveBackticks(word.DefinitionEn)
	word.DefinitionFr = security.RemoveBackticks(word.DefinitionFr)
	word.DefinitionIt = security.RemoveBackticks(word.DefinitionIt)
	word.DefinitionPt = security.RemoveBackticks(word.DefinitionPt)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// FindWord recieves an id, and returns an specific word with that id.
func FindWord(c *gin.Context) {
	var word models.Word
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Word not found!"})
		return
	}

	word.TextDe = security.RemoveBackticks(word.TextDe)
	word.TextEs = security.RemoveBackticks(word.TextEs)
	word.TextEn = security.RemoveBackticks(word.TextEn)
	word.TextFr = security.RemoveBackticks(word.TextFr)
	word.TextIt = security.RemoveBackticks(word.TextIt)
	word.TextPt = security.RemoveBackticks(word.TextPt)
	word.ContextDe = security.RemoveBackticks(word.ContextDe)
	word.ContextEs = security.RemoveBackticks(word.ContextEs)
	word.ContextEn = security.RemoveBackticks(word.ContextEn)
	word.ContextFr = security.RemoveBackticks(word.ContextFr)
	word.ContextIt = security.RemoveBackticks(word.ContextIt)
	word.ContextPt = security.RemoveBackticks(word.ContextPt)
	word.DefinitionDe = security.RemoveBackticks(word.DefinitionDe)
	word.DefinitionEs = security.RemoveBackticks(word.DefinitionEs)
	word.DefinitionEn = security.RemoveBackticks(word.DefinitionEn)
	word.DefinitionFr = security.RemoveBackticks(word.DefinitionFr)
	word.DefinitionIt = security.RemoveBackticks(word.DefinitionIt)
	word.DefinitionPt = security.RemoveBackticks(word.DefinitionPt)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// PatchWord updates a word
func PatchWord(c *gin.Context) {

	// Get model if exist
	var word models.Word
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Word not found!"})
		return
	}

	var input models.UpdateWordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&word).Updates(
		models.Word{
			ID:             word.ID,
			LocaleID:       uint(input.LocaleID),
			WordCategoryID: uint(input.WordCategoryID),
			TextDe:         security.SecureString(input.TextDe),
			TextEs:         security.SecureString(input.TextEs),
			TextEn:         security.SecureString(input.TextEn),
			TextFr:         security.SecureString(input.TextFr),
			TextIt:         security.SecureString(input.TextIt),
			TextPt:         security.SecureString(input.TextPt),
			ContextDe:      security.SecureString(input.ContextDe),
			ContextEs:      security.SecureString(input.ContextEs),
			ContextEn:      security.SecureString(input.ContextEn),
			ContextFr:      security.SecureString(input.ContextFr),
			ContextIt:      security.SecureString(input.ContextIt),
			ContextPt:      security.SecureString(input.ContextPt),
			DefinitionDe:   security.SecureString(input.DefinitionDe),
			DefinitionEs:   security.SecureString(input.DefinitionEs),
			DefinitionEn:   security.SecureString(input.DefinitionEn),
			DefinitionFr:   security.SecureString(input.DefinitionFr),
			DefinitionIt:   security.SecureString(input.DefinitionIt),
			DefinitionPt:   security.SecureString(input.DefinitionPt),
			Modified:       t,
		})

	word.TextDe = security.RemoveBackticks(word.TextDe)
	word.TextEs = security.RemoveBackticks(word.TextEs)
	word.TextEn = security.RemoveBackticks(word.TextEn)
	word.TextFr = security.RemoveBackticks(word.TextFr)
	word.TextIt = security.RemoveBackticks(word.TextIt)
	word.TextPt = security.RemoveBackticks(word.TextPt)
	word.ContextDe = security.RemoveBackticks(word.ContextDe)
	word.ContextEs = security.RemoveBackticks(word.ContextEs)
	word.ContextEn = security.RemoveBackticks(word.ContextEn)
	word.ContextFr = security.RemoveBackticks(word.ContextFr)
	word.ContextIt = security.RemoveBackticks(word.ContextIt)
	word.ContextPt = security.RemoveBackticks(word.ContextPt)
	word.DefinitionDe = security.RemoveBackticks(word.DefinitionDe)
	word.DefinitionEs = security.RemoveBackticks(word.DefinitionEs)
	word.DefinitionEn = security.RemoveBackticks(word.DefinitionEn)
	word.DefinitionFr = security.RemoveBackticks(word.DefinitionFr)
	word.DefinitionIt = security.RemoveBackticks(word.DefinitionIt)
	word.DefinitionPt = security.RemoveBackticks(word.DefinitionPt)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// DeleteWord deletes a word
func DeleteWord(c *gin.Context) {
	// Get model if exist
	var word models.Word
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&word)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
