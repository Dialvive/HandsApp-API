package main

import (
	"API/controllers"
	"API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// SIMPLE TABLES ROUTES /////////////////////////////////////////

	// Routes for countries
	r.GET("/countries", controllers.GetCountries)
	r.POST("/country", controllers.CreateCountry)
	r.GET("/country/:id", controllers.FindCountry)
	r.PATCH("/country/:id", controllers.PatchCountry)
	r.DELETE("/country/:id", controllers.DeleteCountry)

	// Routes for spokenLanguages
	r.GET("/spoken_languages", controllers.GetSpokenLanguages)
	r.POST("/spoken_language", controllers.CreateSpokenLanguage)
	r.GET("/spoken_language/:id", controllers.FindSpokenLanguage)
	r.PATCH("/spoken_language/:id", controllers.PatchSpokenLanguage)
	r.DELETE("/spoken_language/:id", controllers.DeleteSpokenLanguage)

	// Routes for signLanguages
	r.GET("/sign_languages", controllers.GetSignLanguages)
	r.POST("/sign_language", controllers.CreateSignLanguage)
	r.GET("/sign_language/:id", controllers.FindSignLanguage)
	r.PATCH("/sign_language/:id", controllers.PatchSignLanguage)
	r.DELETE("/sign_language/:id", controllers.DeleteSignLanguage)

	// Routes for phraseCategories
	r.GET("/phrase_categories", controllers.GetPhraseCategories)
	r.POST("/phrase_category", controllers.CreatePhraseCategory)
	r.GET("/phrase_category/:id", controllers.FindPhraseCategory)
	r.PATCH("/phrase_category/:id", controllers.PatchPhraseCategory)
	r.DELETE("/phrase_category/:id", controllers.DeletePhraseCategory)

	// Routes for wordCategories
	r.GET("/word_categories", controllers.GetWordCategories)
	r.POST("/word_category", controllers.CreateWordCategory)
	r.GET("/word_category/:id", controllers.FindWordCategory)
	r.PATCH("/word_category/:id", controllers.PatchWordCategory)
	r.DELETE("/word_category/:id", controllers.DeleteWordCategory)

	// Routes for ad_categories
	r.GET("/ad_categories", controllers.GetAdCategories)
	r.POST("/ad_category", controllers.CreateAdCategory)
	r.GET("/ad_category/:id", controllers.FindAdCategory)
	r.PATCH("/ad_category/:id", controllers.PatchAdCategory)
	r.DELETE("/ad_category/:id", controllers.DeleteAdCategory)

	// Routes for friendships
	r.GET("/friendships", controllers.GetFriendships)
	r.POST("/friendship", controllers.CreateFriendship)
	r.GET("/friendship/:id", controllers.FindFriendship)
	r.PATCH("/friendship/:id", controllers.PatchFriendship)
	r.DELETE("/friendship/:id", controllers.DeleteFriendship)

	// RELATED TABLES ROUTES ////////////////////////////////////////

	// Routes for regions
	r.GET("/regions", controllers.GetRegions)
	r.POST("/region", controllers.CreateRegion)
	r.GET("/region/:id", controllers.FindRegion)
	r.PATCH("/region/:id", controllers.PatchRegion)
	r.DELETE("/region/:id", controllers.DeleteRegion)

	// Routes for users
	r.GET("/users", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.FindUser)
	r.PATCH("/user/:id", controllers.PatchUser)
	r.PUT("/user/:id", controllers.PutUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// Routes for locales
	r.GET("/locales", controllers.GetLocales)
	r.POST("/locale", controllers.CreateLocale)
	r.GET("/locale/:id", controllers.FindLocale)
	r.PATCH("/locale/:id", controllers.PatchLocale)
	r.DELETE("/locale/:id", controllers.DeleteLocale)

	// Routes for words
	r.GET("/words", controllers.GetWords)
	r.POST("/word", controllers.CreateWord)
	r.GET("/word/:id", controllers.FindWord)
	r.PATCH("/word/:id", controllers.PatchWord)
	r.DELETE("/word/:id", controllers.DeleteWord)

	// Routes for phrases
	r.GET("/phrases", controllers.GetPhrases)
	r.POST("/phrase", controllers.CreatePhrase)
	r.GET("/phrase/:id", controllers.FindPhrase)
	r.PATCH("/phrase/:id", controllers.PatchPhrase)
	r.DELETE("/phrase/:id", controllers.DeletePhrase)

	// Routes for advertisements
	r.GET("/advertisements", controllers.GetAdvertisements)
	r.POST("/advertisement", controllers.CreateAdvertisement)
	r.GET("/advertisement/:id", controllers.FindAdvertisement)
	r.PATCH("/advertisement/:id", controllers.PatchAdvertisement)
	r.DELETE("/advertisement/:id", controllers.DeleteAdvertisement)

	r.Run(":8080")
}
