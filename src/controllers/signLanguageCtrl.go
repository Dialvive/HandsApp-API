package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSignLanguages retrieves all the signLanguages from the DB.
func GetSignLanguages(c *gin.Context) {
	var signLanguages []models.SignLanguage
	models.DB.Find(&signLanguages)

	for i := range signLanguages {
		signLanguages[i].NameDe = security.RemoveBackticks(signLanguages[i].NameDe)
		signLanguages[i].NameEs = security.RemoveBackticks(signLanguages[i].NameEs)
		signLanguages[i].NameEn = security.RemoveBackticks(signLanguages[i].NameEn)
		signLanguages[i].NameFr = security.RemoveBackticks(signLanguages[i].NameFr)
		signLanguages[i].NameIt = security.RemoveBackticks(signLanguages[i].NameIt)
		signLanguages[i].NamePt = security.RemoveBackticks(signLanguages[i].NamePt)
		signLanguages[i].Abbreviation = security.RemoveBackticks(signLanguages[i].Abbreviation)
	}

	c.JSON(http.StatusOK, gin.H{"data": signLanguages})
}

// CreateSignLanguage creates a new signLanguage.
func CreateSignLanguage(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		time.Sleep(5 * time.Second)
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.CreateSignLanguageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	signLanguage := models.SignLanguage{
		NameDe:       security.SecureString(input.NameDe),
		NameEs:       security.SecureString(input.NameEs),
		NameEn:       security.SecureString(input.NameEn),
		NameFr:       security.SecureString(input.NameFr),
		NameIt:       security.SecureString(input.NameIt),
		NamePt:       security.SecureString(input.NamePt),
		Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 6)),
		Modified:     t}
	models.DB.Create(&signLanguage)

	signLanguage.NameDe = security.RemoveBackticks(signLanguage.NameDe)
	signLanguage.NameEs = security.RemoveBackticks(signLanguage.NameEs)
	signLanguage.NameEn = security.RemoveBackticks(signLanguage.NameEn)
	signLanguage.NameFr = security.RemoveBackticks(signLanguage.NameFr)
	signLanguage.NameIt = security.RemoveBackticks(signLanguage.NameIt)
	signLanguage.NamePt = security.RemoveBackticks(signLanguage.NamePt)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// FindSignLanguage recieves an id, and returns an specific signLanguage with that id.
func FindSignLanguage(c *gin.Context) {
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignLanguage not found!"})
		return
	}

	signLanguage.NameDe = security.RemoveBackticks(signLanguage.NameDe)
	signLanguage.NameEs = security.RemoveBackticks(signLanguage.NameEs)
	signLanguage.NameEn = security.RemoveBackticks(signLanguage.NameEn)
	signLanguage.NameFr = security.RemoveBackticks(signLanguage.NameFr)
	signLanguage.NameIt = security.RemoveBackticks(signLanguage.NameIt)
	signLanguage.NamePt = security.RemoveBackticks(signLanguage.NamePt)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// PatchSignLanguage updates a signLanguage
func PatchSignLanguage(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		time.Sleep(5 * time.Second)
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SignLanguage not found!"})
		return
	}

	var input models.UpdateSignLanguageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&signLanguage).Updates(
		models.SignLanguage{
			ID:           signLanguage.ID,
			NameDe:       security.SecureString(input.NameDe),
			NameEs:       security.SecureString(input.NameEs),
			NameEn:       security.SecureString(input.NameEn),
			NameFr:       security.SecureString(input.NameFr),
			NameIt:       security.SecureString(input.NameIt),
			NamePt:       security.SecureString(input.NamePt),
			Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 2)),
			Modified:     t,
		})

	signLanguage.NameDe = security.RemoveBackticks(signLanguage.NameDe)
	signLanguage.NameEs = security.RemoveBackticks(signLanguage.NameEs)
	signLanguage.NameEn = security.RemoveBackticks(signLanguage.NameEn)
	signLanguage.NameFr = security.RemoveBackticks(signLanguage.NameFr)
	signLanguage.NameIt = security.RemoveBackticks(signLanguage.NameIt)
	signLanguage.NamePt = security.RemoveBackticks(signLanguage.NamePt)
	signLanguage.Abbreviation = security.RemoveBackticks(signLanguage.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": signLanguage})
}

// DeleteSignLanguage deletes a signLanguage
func DeleteSignLanguage(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		time.Sleep(5 * time.Second)
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var signLanguage models.SignLanguage
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&signLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&signLanguage)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
