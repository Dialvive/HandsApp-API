package main

import (
	"API/controllers"
	"API/models"
	"API/security"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//TODO: 1) USE JWT

	//TODO: 2) USE CACHES
	// * implemented caches

	//TODO: 3) ALLOW FILTERS

	//TODO: 4) VALIDATE USER INPUT PRE DB
	//TODO: * RETURN ERROR IF NECESSARY ENTITIES DONT EXIST WHILE CREATING & UPDATING
	//TODO: * DONT ALLOW DELETING RELATED ENTITY IF THERE ARE OBJECTS RELATED TO IT
	//TODO: * DELETING AN USER DELETES EVERYTHING RELATED TO IT

	//TODO: 7) FIX friends count

	// ! PRODUCTION ONLY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	/*
		gin.SetMode(gin.ReleaseMode)

		// GET HTTPS/TLS CERTIFICATES

		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/var/www/.cache"),
			HostPolicy: autocert.HostWhitelist("api.handsapp.org"),
			Email:      "haikode@protonmail.com",
		}

		// REDIRECT ALL HTTP TO HTTPS

		httpRouter := gin.Default()
		httpRouter.NoRoute(func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "https://api.handsapp.org"+c.FullPath())
		})
	*/
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	r := gin.Default()
	models.ConnectDatabase()
	models.ConnectMeili()
	controllers.PopulateMeili()

	// CORS POLICY //////////////////////////////////////////////////

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// EVERY OTHER ROUTE ////////////////////////////////////////////

	r.NoRoute(security.RerouteHandler)
	r.Any("/", security.RerouteHandler)

	// SIMPLE TABLES ROUTES /////////////////////////////////////////

	// Routes for countries
	r.GET("/v1/countries", controllers.GetCountries)
	r.POST("/v1/country", controllers.CreateCountry)
	r.GET("/v1/country/:ID", controllers.FindCountry)
	r.PATCH("/v1/country/:ID", controllers.PatchCountry)
	r.DELETE("/v1/country/:ID", controllers.DeleteCountry)

	// Routes for spokenLanguages
	r.GET("/v1/spoken_languages", controllers.GetSpokenLanguages)
	r.POST("/v1/spoken_language", controllers.CreateSpokenLanguage)
	r.GET("/v1/spoken_language/:ID", controllers.FindSpokenLanguage)
	r.PATCH("/v1/spoken_language/:ID", controllers.PatchSpokenLanguage)
	r.DELETE("/v1/spoken_language/:ID", controllers.DeleteSpokenLanguage)

	// Routes for signLanguages
	r.GET("/v1/sign_languages", controllers.GetSignLanguages)
	r.POST("/v1/sign_language", controllers.CreateSignLanguage)
	r.GET("/v1/sign_language/:ID", controllers.FindSignLanguage)
	r.PATCH("/v1/sign_language/:ID", controllers.PatchSignLanguage)
	r.DELETE("/v1/sign_language/:ID", controllers.DeleteSignLanguage)

	// Routes for phraseCategories
	/*
		r.GET("/v1/phrase_categories", controllers.GetPhraseCategories)
		r.POST("/v1/phrase_category", controllers.CreatePhraseCategory)
		r.GET("/v1/phrase_category/:ID", controllers.FindPhraseCategory)
		r.PATCH("/v1/phrase_category/:ID", controllers.PatchPhraseCategory)
		r.DELETE("/v1/phrase_category/:ID", controllers.DeletePhraseCategory)
	*/

	// Routes for wordCategories
	r.GET("/v1/word_categories", controllers.GetWordCategories)
	r.POST("/v1/word_category", controllers.CreateWordCategory)
	r.GET("/v1/word_category/:ID", controllers.FindWordCategory)
	r.PATCH("/v1/word_category/:ID", controllers.PatchWordCategory)
	r.DELETE("/v1/word_category/:ID", controllers.DeleteWordCategory)

	// Routes for ad_categories
	/*
		r.GET("/v1/ad_categories", controllers.GetAdCategories)
		r.POST("/v1/ad_category", controllers.CreateAdCategory)
		r.GET("/v1/ad_category/:ID", controllers.FindAdCategory)
		r.PATCH("/v1/ad_category/:ID", controllers.PatchAdCategory)
		r.DELETE("/v1/ad_category/:ID", controllers.DeleteAdCategory)

		// Routes for friendships
		r.GET("/v1/friendships", controllers.GetFriendships)
		r.POST("/v1/friendship", controllers.CreateFriendship)
		r.GET("/v1/friendship/:ID", controllers.FindFriendship)
		r.PATCH("/v1/friendship/:ID", controllers.PatchFriendship)
		r.DELETE("/v1/friendship/:ID", controllers.DeleteFriendship)
	*/

	// RELATED TABLES ROUTES ////////////////////////////////////////

	// Routes for regions
	r.GET("/v1/regions", controllers.GetRegions)
	r.POST("/v1/region", controllers.CreateRegion)
	r.GET("/v1/region/:ID", controllers.FindRegion)
	r.PATCH("/v1/region/:ID", controllers.PatchRegion)
	r.DELETE("/v1/region/:ID", controllers.DeleteRegion)

	// Routes for users
	/*
		r.GET("/v1/users", controllers.GetUsers)
		r.POST("/v1/user", controllers.CreateUser)
		r.GET("/v1/user/:ID", controllers.FindUser)
		r.PATCH("/v1/user/:ID", controllers.PatchUser)
		r.PUT("/v1/user/:ID", controllers.PutUser)
		r.DELETE("/v1/user/:ID", controllers.DeleteUser)
	*/

	// Routes for locales
	r.GET("/v1/locales", controllers.GetLocales)
	r.POST("/v1/locale", controllers.CreateLocale)
	r.GET("/v1/locale/:ID", controllers.FindLocale)
	r.PATCH("/v1/locale/:ID", controllers.PatchLocale)
	r.DELETE("/v1/locale/:ID", controllers.DeleteLocale)

	// Routes for words
	r.GET("/v1/words", controllers.GetWords)
	r.POST("/v1/word", controllers.CreateWord)
	r.GET("/v1/word/:ID", controllers.FindWord)
	r.PATCH("/v1/word/:ID", controllers.PatchWord)
	r.DELETE("/v1/word/:ID", controllers.DeleteWord)

	// Routes for phrases
	/*
		r.GET("/v1/phrases", controllers.GetPhrases)
		r.POST("/v1/phrase", controllers.CreatePhrase)
		r.GET("/v1/phrase/:ID", controllers.FindPhrase)
		r.PATCH("/v1/phrase/:ID", controllers.PatchPhrase)
		r.DELETE("/v1/phrase/:ID", controllers.DeletePhrase)

		// Routes for advertisements
		r.GET("/v1/advertisements", controllers.GetAdvertisements)
		r.POST("/v1/advertisement", controllers.CreateAdvertisement)
		r.GET("/v1/advertisement/:ID", controllers.FindAdvertisement)
		r.PATCH("/v1/advertisement/:ID", controllers.PatchAdvertisement)
		r.DELETE("/v1/advertisement/:ID", controllers.DeleteAdvertisement)
	*/

	// Weak entities routes /////////////////////////////////////////

	// Routes for wordSigns

	r.GET("/v1/word_signs", controllers.GetWordSigns)
	r.POST("/v1/word_sign", controllers.CreateWordSign)
	r.GET("/v1/word_signs/:wordID", controllers.FindAllWordSigns)
	r.GET("/v1/word_signs/:wordID/:localeID", controllers.FindWordSigns)
	r.GET("/v1/word_sign/:wordID/:localeID/:version", controllers.FindWordSign)
	r.GET("/v1/word_sign_count/:wordID/:localeID", controllers.CountWordSigns)
	r.DELETE("/v1/word_sign/:wordID/:localeID/:version", controllers.DeleteWordSign)

	// Routes for phraseSigns
	/*
		r.GET("/v1/phrase_signs", controllers.GetPhraseSigns)
		r.POST("/v1/phrase_sign", controllers.CreatePhraseSign)
		r.GET("/v1/phrase_signs/:phraseID", controllers.FindPhraseSigns)
		r.GET("/v1/phrase_sign/:phraseID/:localeID", controllers.FindPhraseSign)
		r.DELETE("/v1/phrase_sign/:phraseID/:localeID/", controllers.DeletePhraseSign)

		// Routes for friends
		r.GET("/v1/friends", controllers.GetFriends)
		r.POST("/v1/friend", controllers.CreateFriend)
		r.GET("/v1/friend/:ID1/:ID2", controllers.FindFriend)
		r.GET("/v1/friends/:ID", controllers.FindFriends)
		r.GET("/v1/friends_count/:ID", controllers.CountFriends)
		r.PATCH("/v1/friend/:ID1/:ID2", controllers.PatchFriend)
		r.DELETE("/v1/friend/:ID1/:ID2", controllers.DeleteFriend)

		// Routes for favorite_phrases
		r.GET("/v1/favorite_phrases", controllers.GetFavoritePhrases)
		r.POST("/v1/favorite_phrase", controllers.CreateFavoritePhrases)
		r.GET("/v1/favorite_phrases_count/phrase/:phraseID", controllers.CountFavoritePhrasesP)
		r.GET("/v1/favorite_phrases_count/user/:userID", controllers.CountFavoritePhrasesU)
		r.GET("/v1/favorite_phrases/:userID", controllers.FindFavoritePhrases)
		r.DELETE("/v1/favorite_phrase/:userID/:phraseID", controllers.DeleteFavoritePhrases)

		// Routes for favorite_words
		r.GET("/v1/favorite_words", controllers.GetFavoriteWords)
		r.POST("/v1/favorite_word", controllers.CreateFavoriteWords)
		r.GET("/v1/favorite_words_count/word/:wordID", controllers.CountFavoriteWordsP)
		r.GET("/v1/favorite_words_count/user/:userID", controllers.CountFavoriteWordsU)
		r.GET("/v1/favorite_words/:userID", controllers.FindFavoriteWords)
		r.DELETE("/v1/favorite_word/:userID/:wordID", controllers.DeleteFavoriteWords)
	*/

	// Routes for Meilisearch
	r.POST("/v1/search/words", controllers.MeiliSearchWords)

	// ! PRODUCTION ONLY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	/*
		log.Fatal(autotls.RunWithManager(r, &m)) // HTTPS
		log.Fatal(httpRouter.Run(":80"))         // HTTP
	*/
	// DEVELOPMENT ONLY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	log.Fatal(r.Run(":8080"))

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}
