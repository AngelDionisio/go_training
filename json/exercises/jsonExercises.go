package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type searchPayload struct {
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

// func Unmarshal(data []byte, v interface{}) error
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
// value to be stored needs to be a pointer to the struct
func unmarshalExercise() {
	jsonString := `{"searchQuery":{"start":0,"num":24,"sortOrder":5,"listingTypes":[2],"agentSearch":false,"geography":"nyc"},"relationTypes":[0]}`
	sbJSON := []byte(jsonString)

	// variable will be passed to unmarshal, as the v interface{} parameter, the results of parsing the json will be stored here
	var storedFromJSON searchPayload

	// json.Unmarshal takes a []bye, and a pointer to a struct that can hold the parsed JSON encoded data
	err := json.Unmarshal(sbJSON, &storedFromJSON)
	if err != nil {
		fmt.Println("error trying to unmarshal JSON", err)
	}

	fmt.Println(storedFromJSON)
	fmt.Printf("%+v\n", storedFromJSON)
}

func jsonMarshallExercise() {
	p1 := person{
		First: "Michael",
		Last:  "Jordan",
		Age:   45,
	}

	p2 := person{
		First: "Scotty",
		Last:  "Pippen",
		Age:   44,
	}

	basketballPlayers := []person{p1, p2}
	// marshall returns a slice of bytes
	// func Marshal(v interface{}) ([]byte, error)
	jsonRepresentation, err := json.Marshal(basketballPlayers)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v\n", string(jsonRepresentation))
	fmt.Println(string(jsonRepresentation))
}

// Marshal
// func Marshal(v interface{}) ([]byte, error)
// returns the JSON encoding of v

// Unmarshal
// func Unmarshal(data []byte. v interface{}) error
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
// helpful site: https://mholt.github.io/json-to-go/

func main() {
	// jsonMarshallExercise()
	unmarshalExercise()
}
