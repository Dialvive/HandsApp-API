package controllers

import (
	"API/models"
	"API/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetCountries retrieves all the countries from the DB.
func GetCountries(c *gin.Context) {
	var countries []models.Country
	models.DB.Find(&countries)

	for i := range countries {
		countries[i].NameDe = security.RemoveBackticks(countries[i].NameDe)
		countries[i].NameEs = security.RemoveBackticks(countries[i].NameEs)
		countries[i].NameEn = security.RemoveBackticks(countries[i].NameEn)
		countries[i].NameFr = security.RemoveBackticks(countries[i].NameFr)
		countries[i].NameIt = security.RemoveBackticks(countries[i].NameIt)
		countries[i].NamePt = security.RemoveBackticks(countries[i].NamePt)
		countries[i].Abbreviation = security.RemoveBackticks(countries[i].Abbreviation)
	}

	c.JSON(http.StatusOK, gin.H{"data": countries})
}

// CreateCountry recieves a Name and Abbreviation, and creates a new country.
func CreateCountry(c *gin.Context) {
	var input models.CreateCountryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	country := models.Country{
		NameDe:       security.SecureString(input.NameDe),
		NameEs:       security.SecureString(input.NameEs),
		NameEn:       security.SecureString(input.NameEn),
		NameFr:       security.SecureString(input.NameFr),
		NameIt:       security.SecureString(input.NameIt),
		NamePt:       security.SecureString(input.NamePt),
		Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 2)),
		Modified:     t}
	models.DB.Create(&country)

	country.NameDe = security.RemoveBackticks(country.NameDe)
	country.NameEs = security.RemoveBackticks(country.NameEs)
	country.NameEn = security.RemoveBackticks(country.NameEn)
	country.NameFr = security.RemoveBackticks(country.NameFr)
	country.NameIt = security.RemoveBackticks(country.NameIt)
	country.NamePt = security.RemoveBackticks(country.NamePt)
	country.Abbreviation = security.RemoveBackticks(country.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// FindCountry recieves an id, and returns an specific country with that id.
func FindCountry(c *gin.Context) {
	var country models.Country
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country not found!"})
		return
	}

	country.NameDe = security.RemoveBackticks(country.NameDe)
	country.NameEs = security.RemoveBackticks(country.NameEs)
	country.NameEn = security.RemoveBackticks(country.NameEn)
	country.NameFr = security.RemoveBackticks(country.NameFr)
	country.NameIt = security.RemoveBackticks(country.NameIt)
	country.NamePt = security.RemoveBackticks(country.NamePt)
	country.Abbreviation = security.RemoveBackticks(country.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// PatchCountry updates a country
func PatchCountry(c *gin.Context) {

	// Get model if exist
	var country models.Country
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country not found!"})
		return
	}

	var input models.UpdateCountryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&country).Updates(
		models.Country{
			ID:           country.ID,
			NameDe:       security.SecureString(input.NameDe),
			NameEs:       security.SecureString(input.NameEs),
			NameEn:       security.SecureString(input.NameEn),
			NameFr:       security.SecureString(input.NameFr),
			NameIt:       security.SecureString(input.NameIt),
			NamePt:       security.SecureString(input.NamePt),
			Abbreviation: security.SecureString(security.TrimToLength(input.Abbreviation, 2)),
			Modified:     t,
		})

	country.NameDe = security.RemoveBackticks(country.NameDe)
	country.NameEs = security.RemoveBackticks(country.NameEs)
	country.NameEn = security.RemoveBackticks(country.NameEn)
	country.NameFr = security.RemoveBackticks(country.NameFr)
	country.NameIt = security.RemoveBackticks(country.NameIt)
	country.NamePt = security.RemoveBackticks(country.NamePt)
	country.Abbreviation = security.RemoveBackticks(country.Abbreviation)

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// DeleteCountry deletes a country
func DeleteCountry(c *gin.Context) {
	// Get model if exist
	var country models.Country
	var param uint64
	var err error
	if param, err = security.SecureUint(c.Param("ID")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := models.DB.Where("id = ?", param).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&country)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
