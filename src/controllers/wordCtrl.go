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
		words[i].Context = security.RemoveBackticks(words[i].Context)
		words[i].Definition = security.RemoveBackticks(words[i].Definition)
		words[i].Text = security.RemoveBackticks(words[i].Text)
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
		Text:           security.SecureString(input.Text),
		Definition:     security.SecureString(input.Definition),
		Context:        security.SecureString(input.Context),
		Modified:       t,
	}
	models.DB.Create(&word)

	word.Context = security.RemoveBackticks(word.Context)
	word.Definition = security.RemoveBackticks(word.Definition)
	word.Text = security.RemoveBackticks(word.Text)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// FindWord recieves an id, and returns an specific word with that id.
func FindWord(c *gin.Context) {
	var word models.Word

	if err := models.DB.Where("id = ?", c.Param("id")).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Word not found!"})
		return
	}

	word.Context = security.RemoveBackticks(word.Context)
	word.Definition = security.RemoveBackticks(word.Definition)
	word.Text = security.RemoveBackticks(word.Text)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// PatchWord updates a word
func PatchWord(c *gin.Context) {

	// Get model if exist
	var word models.Word

	if err := models.DB.Where("id = ?", c.Param("id")).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Word not found!"})
		return
	}

	var input models.CreateWordInput

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
			Text:           security.SecureString(input.Text),
			Definition:     security.SecureString(input.Definition),
			Context:        security.SecureString(input.Context),
			Modified:       t,
		})

	word.Context = security.RemoveBackticks(word.Context)
	word.Definition = security.RemoveBackticks(word.Definition)
	word.Text = security.RemoveBackticks(word.Text)

	c.JSON(http.StatusOK, gin.H{"data": word})
}

// DeleteWord deletes a word
func DeleteWord(c *gin.Context) {
	// Get model if exist
	var word models.Word
	if err := models.DB.Where("id = ?", c.Param("id")).First(&word).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&word)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
