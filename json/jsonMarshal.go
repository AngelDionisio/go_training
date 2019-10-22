package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Address string
}

// Marshal
// func Marshal(v interface{}) ([]byte, error)
// returns the JSON encoding of v

// Unmarshal
// func Unmarshal(data []byte. v interface{}) error
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
// helpful site: https://mholt.github.io/json-to-go/

func main() {
	// p1 equal to composite literal person
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	// func Marshal(v interface{}) ([]byte, error)
	sliceOfBytes, err := json.Marshal(people)

	if err != nil {
		fmt.Println("there was an error trying to convert to JSON", err)
	}

	fmt.Println("JSON representation of people", string(sliceOfBytes))

	usingUnmarshal()

}

func usingUnmarshal() {
	var jsonblob = `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	var sbBlob = []byte(jsonblob)

	fmt.Printf("%T\n", jsonblob) // string
	fmt.Printf("%T\n", sbBlob)   // []uint8

	var people []person

	// Unmarshal returns and error, if successful, it will write to the address of the pointer passed
	// func Unmarshal(data []byte, v interface{}) error
	err := json.Unmarshal(sbBlob, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	// %v prints value in default format, when printing structs adding the plug flag adds fields names
	fmt.Printf("%+v\n", people)
}
