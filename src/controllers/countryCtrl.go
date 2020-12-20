package controllers

import (
	"API/models"
	"net/http"

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

	country := models.Country{Name: input.Name, Abbreviation: input.Abbreviation}
	models.DB.Create(&country)

	c.JSON(http.StatusOK, gin.H{"data": country})
}
