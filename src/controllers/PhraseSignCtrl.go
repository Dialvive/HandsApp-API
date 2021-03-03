package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPhraseSigns retrieves all the phraseSigns from the DB.
func GetPhraseSigns(c *gin.Context) {
	var phraseSigns []models.PhraseSign
	models.DB.Find(&phraseSigns)

	c.JSON(http.StatusOK, gin.H{"data": phraseSigns})
}

// CreatePhraseSign creates a new phraseSigns.
func CreatePhraseSign(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.PhraseSignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phraseSign := models.PhraseSign{
		PhraseID: input.PhraseID,
		LocaleID: input.LocaleID,
		RegionID: input.RegionID,
		Modified: time.Now().UTC().Format("2006-01-02 15:04:05"),
	}
	models.DB.Create(&phraseSign)

	c.JSON(http.StatusOK, gin.H{"data": phraseSign})
}

// FindPhraseSigns recieves a Phrase id, and returns all its signs.
func FindPhraseSigns(c *gin.Context) {
	var signs []models.PhraseSign
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("phraseID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := models.DB.Where("phrase_ID = ?", param).Find(&signs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signs not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": signs})
}

// FindPhraseSign recieves a PhraseID and LocaleID, and returns all of the signs with the same phrase and locale.
func FindPhraseSign(c *gin.Context) {
	var signs []models.PhraseSign
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("phraseID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := models.DB.Where("phrase_ID = ? AND locale_ID = ?", param1, param2).Find(&signs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signs not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": signs})
}

// DeletePhraseSign deletes a phraseSign
func DeletePhraseSign(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var sign models.PhraseSign
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("phraseID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	// En lugar de Delete debería ser First, pero no funciona así ---------------VVVVVV
	if err := models.DB.Where("phrase_ID = ? AND locale_ID = ?", param1, param2).Delete(&sign).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//models.DB.Delete(&sign) Porqué no funciona ésto?
	c.JSON(http.StatusOK, gin.H{"data": true})
}
