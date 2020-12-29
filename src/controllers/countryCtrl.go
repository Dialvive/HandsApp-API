package controllers

import (
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetCountries retrieves all the countries from the DB.
func GetCountries(c *gin.Context) {
	var countries []models.Country
	models.DB.Find(&countries)

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
		Name:         input.Name,
		Abbreviation: input.Abbreviation,
		Creation:     t}
	models.DB.Create(&country)

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// FindCountry recieves an id, and returns an specific country with that id.
func FindCountry(c *gin.Context) {
	var country models.Country

	if err := models.DB.Where("id = ?", c.Param("id")).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// PatchCountry updates a country
func PatchCountry(c *gin.Context) {

	// Get model if exist
	var country models.Country

	if err := models.DB.Where("id = ?", c.Param("id")).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country not found!"})
		return
	}

	var input models.CreateCountryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	models.DB.Model(&country).Updates(
		models.Country{
			ID:           country.ID,
			Name:         input.Name,
			Abbreviation: input.Abbreviation,
			Creation:     t,
		})

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// DeleteCountry deletes a country
func DeleteCountry(c *gin.Context) {
	// Get model if exist
	var country models.Country
	if err := models.DB.Where("id = ?", c.Param("id")).First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&country)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
