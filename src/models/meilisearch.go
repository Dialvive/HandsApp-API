package models

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

// Meili is a Meilisearch Client
var Meili meilisearch.ClientInterface

// ConnectMeili connects the API to the Meilisearch Client
func ConnectMeili() {
	Meili = meilisearch.NewClient(meilisearch.Config{
		Host: "http://127.0.0.1:7700",
	})
	s, _ := Meili.Stats().GetAll()
	k, _ := Meili.Keys().Get()
	log.Println("Meilisearch: databaseSize: ", s.DatabaseSize)
	log.Println("Meilisearch: lastUpdate:", s.LastUpdate)
	log.Println("Meilisearch: indexes:", s.Indexes)
	log.Println("Meilisearch: public key:", k.Public)
}
