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

	// Routes for signLanguages
	r.GET("/sign_languages", controllers.GetSignLanguages)
	r.POST("/sign_language", controllers.CreateSignLanguage)
	r.GET("/sign_language/:id", controllers.FindSignLanguage)
	r.PATCH("/sign_language/:id", controllers.UpdateSignLanguage)
	r.DELETE("/sign_language/:id", controllers.DeleteSignLanguage)

	// Routes for phraseCategories
	r.GET("/phrase_categories", controllers.GetPhraseCategories)
	r.POST("/phrase_category", controllers.CreatePhraseCategory)
	r.GET("/phrase_category/:id", controllers.FindPhraseCategory)
	r.PATCH("/phrase_category/:id", controllers.UpdatePhraseCategory)
	r.DELETE("/phrase_category/:id", controllers.DeletePhraseCategory)

	// Routes for wordCategories
	r.GET("/word_categories", controllers.GetWordCategories)
	r.POST("/word_category", controllers.CreateWordCategory)
	r.GET("/word_category/:id", controllers.FindWordCategory)
	r.PATCH("/word_category/:id", controllers.UpdateWordCategory)
	r.DELETE("/word_category/:id", controllers.DeleteWordCategory)

	r.Run(":8080")
}
