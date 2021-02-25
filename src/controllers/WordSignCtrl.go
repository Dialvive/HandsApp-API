package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWordSigns retrieves all the wordSigns from the DB.
func GetWordSigns(c *gin.Context) {
	var wordSigns []models.WordSign
	models.DB.Find(&wordSigns)

	c.JSON(http.StatusOK, gin.H{"data": wordSigns})
}

// CreateWordSign creates a new wordSigns.
func CreateWordSign(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		time.Sleep(5 * time.Second)
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.WordSignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wordSign := models.WordSign{
		WordID:   input.WordID,
		LocaleID: input.LocaleID,
		Version:  string(security.SecureString(input.Version)[1]),
		RegionID: input.RegionID,
		Modified: time.Now().UTC().Format("2006-01-02 15:04:05"),
	}
	models.DB.Create(&wordSign)

	c.JSON(http.StatusOK, gin.H{"data": wordSign})
}

// FindAllWordSigns recieves a Word id, and returns all its signs.
func FindAllWordSigns(c *gin.Context) {
	var signs []models.WordSign
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("wordID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := models.DB.Where("word_ID = ?", param).Find(&signs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signs not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": signs})
}

// FindWordSigns recieves a WordID and LocaleID, and returns all of the signs with the same word and locale.
func FindWordSigns(c *gin.Context) {
	var signs []models.WordSign
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("wordID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := models.DB.Where("word_ID = ? AND locale_ID = ?", param1, param2).Find(&signs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signs not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": signs})
}

// FindWordSign recieves a WordID, LocaleID and version, and returns a unique sign with the same parameters if existx.
func FindWordSign(c *gin.Context) {
	var sign models.WordSign
	var param1, param2 uint64
	var param3 string
	var err error
	if param1, err = security.SecureUint(c.Param("wordID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param3 = security.SecureString(c.Param("version")); !strings.ContainsAny(param3, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version"})
		return
	}
	if err := models.DB.Where("word_ID = ? AND locale_ID = ? AND version = ?", param1, param2, string(param3[1])).First(&sign).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sign not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sign})
}

// CountWordSigns recieves a wordID and localeID, returns the number of versions that wordSign has.
// it as favorite.
func CountWordSigns(c *gin.Context) {
	var count int64
	var param1, param2 uint64
	var err error
	if param1, err = security.SecureUint(c.Param("wordID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := models.DB.Model(&models.WordSign{}).
		Where("word_ID = ? AND locale_ID = ?", param1, param2).Count(&count).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// DeleteWordSign deletes a wordSign
func DeleteWordSign(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		time.Sleep(5 * time.Second)
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var sign models.WordSign
	var param1, param2 uint64
	var param3 string
	var err error
	if param1, err = security.SecureUint(c.Param("wordID")); err != nil || param1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param2, err = security.SecureUint(c.Param("localeID")); err != nil || param2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if param3 = security.SecureString(c.Param("version")); !strings.ContainsAny(param3, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version"})
		return
	}
	if err := models.DB.Where("word_ID = ? AND locale_ID = ? AND version = ?", param1, param2, string(param3[1])).Delete(&sign).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//models.DB.Delete(&sign)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
