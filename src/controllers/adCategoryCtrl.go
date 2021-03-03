package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAdCategories retrieves all the adCategories from the DB.
func GetAdCategories(c *gin.Context) {
	var adCategories []models.AdCategory
	models.DB.Find(&adCategories)

	for i := range adCategories {
		adCategories[i].NameDe = security.RemoveBackticks(adCategories[i].NameDe)
		adCategories[i].NameEs = security.RemoveBackticks(adCategories[i].NameEs)
		adCategories[i].NameEn = security.RemoveBackticks(adCategories[i].NameEn)
		adCategories[i].NameFr = security.RemoveBackticks(adCategories[i].NameFr)
		adCategories[i].NameIt = security.RemoveBackticks(adCategories[i].NameIt)
		adCategories[i].NamePt = security.RemoveBackticks(adCategories[i].NamePt)
	}

	c.JSON(http.StatusOK, gin.H{"data": adCategories})
}

// CreateAdCategory creates a new adCategory.
func CreateAdCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	var input models.CreateAdCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	adCategory := models.AdCategory{
		NameDe:   security.SecureString(input.NameDe),
		NameEs:   security.SecureString(input.NameEs),
		NameEn:   security.SecureString(input.NameEn),
		NameFr:   security.SecureString(input.NameFr),
		NameIt:   security.SecureString(input.NameIt),
		NamePt:   security.SecureString(input.NamePt),
		Cost:     uint(input.Cost),
		Modified: t,
	}
	models.DB.Create(&adCategory)

	adCategory.NameDe = security.RemoveBackticks(adCategory.NameDe)
	adCategory.NameEs = security.RemoveBackticks(adCategory.NameEs)
	adCategory.NameEn = security.RemoveBackticks(adCategory.NameEn)
	adCategory.NameFr = security.RemoveBackticks(adCategory.NameFr)
	adCategory.NameIt = security.RemoveBackticks(adCategory.NameIt)
	adCategory.NamePt = security.RemoveBackticks(adCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// FindAdCategory recieves an id, and returns an specific adCategory with that id.
func FindAdCategory(c *gin.Context) {
	var adCategory models.AdCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	if err := models.DB.Where("id = ?", param).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AdCategory not found!"})
		return
	}

	adCategory.NameDe = security.RemoveBackticks(adCategory.NameDe)
	adCategory.NameEs = security.RemoveBackticks(adCategory.NameEs)
	adCategory.NameEn = security.RemoveBackticks(adCategory.NameEn)
	adCategory.NameFr = security.RemoveBackticks(adCategory.NameFr)
	adCategory.NameIt = security.RemoveBackticks(adCategory.NameIt)
	adCategory.NamePt = security.RemoveBackticks(adCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// PatchAdCategory updates a adCategory
func PatchAdCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var adCategory models.AdCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	if err := models.DB.Where("id = ?", param).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AdCategory not found!"})
		return
	}

	var input models.UpdateAdCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&adCategory).Updates(
		models.AdCategory{
			ID:       adCategory.ID,
			NameDe:   security.SecureString(input.NameDe),
			NameEs:   security.SecureString(input.NameEs),
			NameEn:   security.SecureString(input.NameEn),
			NameFr:   security.SecureString(input.NameFr),
			NameIt:   security.SecureString(input.NameIt),
			NamePt:   security.SecureString(input.NamePt),
			Cost:     uint(input.Cost),
			Modified: t,
		})

	adCategory.NameDe = security.RemoveBackticks(adCategory.NameDe)
	adCategory.NameEs = security.RemoveBackticks(adCategory.NameEs)
	adCategory.NameEn = security.RemoveBackticks(adCategory.NameEn)
	adCategory.NameFr = security.RemoveBackticks(adCategory.NameFr)
	adCategory.NameIt = security.RemoveBackticks(adCategory.NameIt)
	adCategory.NamePt = security.RemoveBackticks(adCategory.NamePt)

	c.JSON(http.StatusOK, gin.H{"data": adCategory})
}

// DeleteAdCategory deletes a adCategory
func DeleteAdCategory(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	// Get model if exist
	var adCategory models.AdCategory
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil || param == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&adCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&adCategory)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
