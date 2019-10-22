package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type geoSearchQuery struct {
	SearchQuery struct {
		Start        int    `json:"start"`
		Num          int    `json:"num"`
		SortOrder    int    `json:"sortOrder"`
		ListingTypes []int  `json:"listingTypes"`
		AgentSearch  bool   `json:"agentSearch"`
		Geography    string `json:"geography"`
	} `json:"searchQuery"`
	RelationTypes []int `json:"relationTypes"`
}

func main() {
	payloadBlob := `{
		"searchQuery": {
		  "start": 0,
		  "num": 24,
		  "sortOrder": 5,
		  "listingTypes": [2],
		  "agentSearch": false,
		  "geography": "nyc"
		},
		"relationTypes": [0]
	}`

	payload := []byte(payloadBlob)

	uri := "https://www.compass.com"
	adapter := "/api/v3/listings/search/list/relations"
	url := uri + adapter

	fmt.Println("THIS IS MY URL:", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("error making http/post request", err)
	}

	var result = make(map[string]interface{})

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["listingRelations"])

	fmt.Println("about to exit")
}
