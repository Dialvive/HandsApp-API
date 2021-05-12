package controllers

import (
	"API/models"
	"API/security"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
)

// PopulateMeili retrieves all the Jsons form the DB and populates the Meilisearch Client
func PopulateMeili(c *gin.Context) {
	if !security.CheckKey(c, c.GetHeader("x-api-key")) && c != nil {
		c.Abort()
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	if words := GetWordsJsonMap(); len(words) == 0 {
		log.Fatal("Couldn't populate meili: Empty set of words")
	} else if _, err := models.Meili.Documents("words").AddOrReplace(words); err != nil {
		log.Fatal("Couldn't populate meili with words: " + err.Error())
	}
}

//MeiliSearchWords returns search results
func MeiliSearchWords(c *gin.Context) {
	var input meilisearch.SearchRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lim int64
	if input.Limit > 100 {
		lim = 10
	} else if input.Limit >= 0 && input.Limit <= 100 {
		lim = input.Limit
	}

	if m, err := models.Meili.Search("words").Search(meilisearch.SearchRequest{
		Query:                 input.Query,
		Offset:                input.Offset,
		Limit:                 lim,
		AttributesToRetrieve:  input.AttributesToRetrieve,
		AttributesToCrop:      input.AttributesToCrop,
		CropLength:            input.CropLength,
		AttributesToHighlight: input.AttributesToHighlight,
		Filters:               input.Filters,
		Matches:               input.Matches,
		FacetsDistribution:    input.FacetsDistribution,
		FacetFilters:          input.FacetFilters,
		PlaceholderSearch:     input.PlaceholderSearch,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, m)
	}

}
