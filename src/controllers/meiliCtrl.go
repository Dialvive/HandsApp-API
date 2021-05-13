package controllers

import (
	"API/models"
	"API/security"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
)

// PopulateMeili retrieves all the Jsons form the DB and populates the Meilisearch Client
func RefreshMeili(c *gin.Context) {
	if c != nil {
		if !security.CheckKey(c, c.GetHeader("x-api-key")) {
			c.Abort()
			c.String(http.StatusNotFound, "404 page not found")
			return
		}
		if err := PopulateMeili(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			s, _ := models.Meili.Stats().GetAll()
			var str string = "databaseSize: " + fmt.Sprint(s.DatabaseSize) + ", lastUpdate:" + s.LastUpdate.String()
			log.Println(str)
			c.JSON(http.StatusOK, str)
		}
	}
}

// PopulateMeili retrieves all the Jsons form the DB and populates the Meilisearch Client
func PopulateMeili() error {
	if words := GetWordsJsonMap(); len(words) == 0 {
		errors.New("Couldn't populate meili: Empty set of words")
	} else if _, err := models.Meili.Documents("words").AddOrReplace(words); err != nil {
		errors.New("Couldn't populate meili with words: " + err.Error())
		return err
	}
	return nil
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
