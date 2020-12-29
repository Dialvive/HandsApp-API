package main

import (
	"API/controllers"
	"API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// SIMPLE TABLES ROUTES

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

	// Routes for spokenLanguages
	r.GET("/sign_languages", controllers.GetSignLanguages)
	r.POST("/sign_language", controllers.CreateSignLanguage)
	r.GET("/sign_language/:id", controllers.FindSignLanguage)
	r.PATCH("/sign_language/:id", controllers.UpdateSignLanguage)
	r.DELETE("/sign_language/:id", controllers.DeleteSignLanguage)

	r.Run(":8080")
}
