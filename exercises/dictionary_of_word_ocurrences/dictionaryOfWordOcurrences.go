package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

var sentence = `We have also come to this hallowed spot to remind America of the fierce urgency of now. 
This is no time to engage in the luxury of cooling off or to take the tranquilizing drug of gradualism. Now is 
the time to make real the promises of democracy; now is the time to rise from the dark and desolate valley of 
segregation to the sunlit path of racial justice; now is the time to lift our nation from the quicksands of racial 
injustice to the solid rock of brotherhood; now is the time to make justice a reality for all of God’s children. It 
would be fatal for the nation to overlook the urgency of the moment. This sweltering summer of the Negro’s legitimate 
discontent will not pass until there is an invigorating autumn of freedom and equality. Nineteen sixty-three is not an end, 
but a beginning. And those who hope that the Negro needed to blow off steam and will now be content, will have a rude awakening if 
the nation returns to business as usual. There will be neither rest nor tranquility in America until the Negro is granted his 
citizenship rights. The whirlwinds of revolt will continue to shake the foundations of our nation until the bright day of justice emerges.`

func main() {
	myMap := CountWordsOcurrences(sentence)
	fmt.Println(myMap)

	fmt.Println("about to exit program", reflect.TypeOf(myMap))
}

// CountWordsOcurrences returns a map words and how many times it appares in the string
// it removes special characters from strings before counting
func CountWordsOcurrences(s string) map[string]int {
	// dict will hold dictionary of ocurrences of each letter
	var dict = make(map[string]int)

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// separate strings by spaces
	sliceOfSeparatedStrings := strings.Split(sentence, " ")
	// sliceOfSeparatedStrings2 := strings.Fields(sentence)

	// iterate over slice of words, clean each one, and then add to dictionary
	// adding one to the default. Since in golang when creating a new var, it takes
	// it's zero value, the value of dict[cleanedStr] on any first ocurrence will be zero
	for _, str := range sliceOfSeparatedStrings {
		cleanedStr := reg.ReplaceAllString(str, "")
		dict[cleanedStr]++
		// dict[cleanedStr] = dict[cleanedStr] + 1
	}

	return dict
}

// WordCount returns a map words and how many times it appares in the string
// it does not clean special characters from strings
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
