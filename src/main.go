package main

import (
	"API/controllers"
	"API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//TODO: 1) USE SSL/TSL
	//TODO: 2) USE JWT
	//TODO: 3) PREVENT SQL INJECTION STRINGS
	//TODO: 4) USE CACHES
	//TODO: 5) ALLOW FILTERS
	// ? USE oAuth?

	r := gin.Default()

	models.ConnectDatabase()

	// SIMPLE TABLES ROUTES /////////////////////////////////////////

	// Routes for countries
	r.GET("/v1/countries", controllers.GetCountries)
	r.POST("/v1/country", controllers.CreateCountry)
	r.GET("/v1/country/:id", controllers.FindCountry)
	r.PATCH("/v1/country/:id", controllers.PatchCountry)
	r.DELETE("/v1/country/:id", controllers.DeleteCountry)

	// Routes for spokenLanguages
	r.GET("/v1/spoken_languages", controllers.GetSpokenLanguages)
	r.POST("/v1/spoken_language", controllers.CreateSpokenLanguage)
	r.GET("/v1/spoken_language/:id", controllers.FindSpokenLanguage)
	r.PATCH("/v1/spoken_language/:id", controllers.PatchSpokenLanguage)
	r.DELETE("/v1/spoken_language/:id", controllers.DeleteSpokenLanguage)

	// Routes for signLanguages
	r.GET("/v1/sign_languages", controllers.GetSignLanguages)
	r.POST("/v1/sign_language", controllers.CreateSignLanguage)
	r.GET("/v1/sign_language/:id", controllers.FindSignLanguage)
	r.PATCH("/v1/sign_language/:id", controllers.PatchSignLanguage)
	r.DELETE("/v1/sign_language/:id", controllers.DeleteSignLanguage)

	// Routes for phraseCategories
	r.GET("/v1/phrase_categories", controllers.GetPhraseCategories)
	r.POST("/v1/phrase_category", controllers.CreatePhraseCategory)
	r.GET("/v1/phrase_category/:id", controllers.FindPhraseCategory)
	r.PATCH("/v1/phrase_category/:id", controllers.PatchPhraseCategory)
	r.DELETE("/v1/phrase_category/:id", controllers.DeletePhraseCategory)

	// Routes for wordCategories
	r.GET("/v1/word_categories", controllers.GetWordCategories)
	r.POST("/v1/word_category", controllers.CreateWordCategory)
	r.GET("/v1/word_category/:id", controllers.FindWordCategory)
	r.PATCH("/v1/word_category/:id", controllers.PatchWordCategory)
	r.DELETE("/v1/word_category/:id", controllers.DeleteWordCategory)

	// Routes for ad_categories
	r.GET("/v1/ad_categories", controllers.GetAdCategories)
	r.POST("/v1/ad_category", controllers.CreateAdCategory)
	r.GET("/v1/ad_category/:id", controllers.FindAdCategory)
	r.PATCH("/v1/ad_category/:id", controllers.PatchAdCategory)
	r.DELETE("/v1/ad_category/:id", controllers.DeleteAdCategory)

	// Routes for friendships
	r.GET("/v1/friendships", controllers.GetFriendships)
	r.POST("/v1/friendship", controllers.CreateFriendship)
	r.GET("/v1/friendship/:id", controllers.FindFriendship)
	r.PATCH("/v1/friendship/:id", controllers.PatchFriendship)
	r.DELETE("/v1/friendship/:id", controllers.DeleteFriendship)

	// RELATED TABLES ROUTES ////////////////////////////////////////

	// Routes for regions
	r.GET("/v1/regions", controllers.GetRegions)
	r.POST("/v1/region", controllers.CreateRegion)
	r.GET("/v1/region/:id", controllers.FindRegion)
	r.PATCH("/v1/region/:id", controllers.PatchRegion)
	r.DELETE("/v1/region/:id", controllers.DeleteRegion)

	// Routes for users
	r.GET("/v1/users", controllers.GetUsers)
	r.POST("/v1/user", controllers.CreateUser)
	r.GET("/v1/user/:id", controllers.FindUser)
	r.PATCH("/v1/user/:id", controllers.PatchUser)
	r.PUT("/v1/user/:id", controllers.PutUser)
	r.DELETE("/v1/user/:id", controllers.DeleteUser)

	// Routes for locales
	r.GET("/v1/locales", controllers.GetLocales)
	r.POST("/v1/locale", controllers.CreateLocale)
	r.GET("/v1/locale/:id", controllers.FindLocale)
	r.PATCH("/v1/locale/:id", controllers.PatchLocale)
	r.DELETE("/v1/locale/:id", controllers.DeleteLocale)

	// Routes for words
	r.GET("/v1/words", controllers.GetWords)
	r.POST("/v1/word", controllers.CreateWord)
	r.GET("/v1/word/:id", controllers.FindWord)
	r.PATCH("/v1/word/:id", controllers.PatchWord)
	r.DELETE("/v1/word/:id", controllers.DeleteWord)

	// Routes for phrases
	r.GET("/v1/phrases", controllers.GetPhrases)
	r.POST("/v1/phrase", controllers.CreatePhrase)
	r.GET("/v1/phrase/:id", controllers.FindPhrase)
	r.PATCH("/v1/phrase/:id", controllers.PatchPhrase)
	r.DELETE("/v1/phrase/:id", controllers.DeletePhrase)

	// Routes for advertisements
	r.GET("/v1/advertisements", controllers.GetAdvertisements)
	r.POST("/v1/advertisement", controllers.CreateAdvertisement)
	r.GET("/v1/advertisement/:id", controllers.FindAdvertisement)
	r.PATCH("/v1/advertisement/:id", controllers.PatchAdvertisement)
	r.DELETE("/v1/advertisement/:id", controllers.DeleteAdvertisement)

	// Weak entities routes /////////////////////////////////////////
	//TODO: FIX FIND METHODS IN WEAK ENTITIES TAKING INTO ACCOUNT TWO IDs
	//TODO: IMPLEMENT FIND ALL WEAK ENTITIES RELATED TO ONE ID

	// Routes for friends
	r.GET("/v1/friends", controllers.GetFriends)
	r.POST("/v1/friend", controllers.CreateFriend)
	r.GET("/v1/friend?:id1&:id2", controllers.FindFriend)
	r.GET("/v1/friends/:id", controllers.FindFriends)
	r.PUT("/v1/friend?:id1&:id2", controllers.PutFriend)
	r.DELETE("/v1/friend?:id1&:id2", controllers.DeleteFriend)

	// Routes for wordsByRegions
	r.GET("/v1/words_by_regions", controllers.GetWordsByRegions)
	r.POST("/v1/word_by_region", controllers.CreateWordByRegion)
	r.GET("/v1/word_by_region?:id1&:id2", controllers.FindWordByRegion)
	r.GET("/v1/words_of_region/:id", controllers.FindWordsOfRegion)
	r.PUT("/v1/word_by_region?:id1&:id2", controllers.PutWordByRegion)
	r.DELETE("/v1/word_by_region?:id1&:id2", controllers.DeleteWordByRegion)

	// Routes for favorite_phrases
	r.GET("/v1/favorite_phrases", controllers.GetFavoritePhrases)
	r.POST("/v1/favorite_phrase", controllers.CreateFavoritePhrase)
	r.GET("/v1/favorite_phrase?:id1&:id2", controllers.FindFavoritePhrase)
	r.GET("/v1/favorite_phrases/:id", controllers.FindFavoritePhrases)
	r.PUT("/v1/favorite_phrase?:id1&:id2", controllers.PutFavoritePhrase)
	r.DELETE("/v1/favorite_phrase?:id1&:id2", controllers.DeleteFavoritePhrase)

	// Routes for favorite_words
	r.GET("/v1/favorite_words", controllers.GetFavoriteWords)
	r.POST("/v1/favorite_word", controllers.CreateFavoriteWord)
	r.GET("/v1/favorite_word?:id1&:id2", controllers.FindFavoriteWord)
	r.GET("/v1/favorite_words/:id", controllers.FindFavoriteWords)
	r.PUT("/v1/favorite_word?:id1&:id2", controllers.PutFavoriteWord)
	r.DELETE("/v1/favorite_word?:id1&:id2", controllers.DeleteFavoriteWord)

	r.Run(":8080")
}
