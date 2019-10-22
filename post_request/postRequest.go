package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type searchPostRequest struct {
	SearchQuery struct {
		Start        int    `json:"start"`
		Num          int    `json:"num"`
		SortOrder    int    `json:"sortOrder"`
		ListingTypes []int  `json:"listingTypes"`
		AgentSearch  bool   `json:"agentSearch"`
		SaleStatuses []int  `json:"saleStatuses"`
		Geography    string `json:"geography"`
	} `json:"searchQuery"`
	RelationTypes []int `json:"relationTypes"`
}

func main() {
	jsonBlob := `{"searchQuery":{"start":0,"num":24,"sortOrder":90,"listingTypes":[2],"agentSearch":true,"saleStatuses":[9,12],"geography":"nyc"},"relationTypes":[0]}`
	bytss := []byte(jsonBlob)

	resp, err := http.Post("https://www.compass.com/api/v3/listings/search/list/relations", "application/json", bytes.NewBuffer(bytss))
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("***************************************************************************")
	fmt.Println(resp)
	fmt.Println("***************************************************************************")

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
