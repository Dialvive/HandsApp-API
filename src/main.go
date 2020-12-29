package main

import (
	"API/controllers"
	"API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Routes for countries
	r.GET("/countries", controllers.GetCountries)
	r.POST("/country", controllers.CreateCountry)
	r.GET("/country/:id", controllers.FindCountry)
	r.PATCH("/country/:id", controllers.UpdateCountry)
	r.DELETE("/country/:id", controllers.DeleteCountry)

	// Routes for spokenLanguages
	r.GET("/spoken_languages", controllers.GetSpokenLanguages)
	r.POST("/spoken_language", controllers.CreateSpokenLanguage)
	r.GET("/spoken_language/:id", controllers.FindSpokenLanguage)
	r.PATCH("/spoken_language/:id", controllers.UpdateSpokenLanguage)
	r.DELETE("/spoken_language/:id", controllers.DeleteSpokenLanguage)

	r.Run(":8080")
}
