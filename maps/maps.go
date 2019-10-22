package main

import (
	"fmt"
)

// maps are key value stores that are optimized for lookups by key
// maps are un-ordered (for range on a map and the results may be different each run)
func main() {
	// createAndQueryMap()
	// makeMap()
	// result := offset("ESTs")
	// fmt.Println(result)

	// forRangeOnMap()
	commaOkIdiomDelete()
}

func commaOkIdiomDelete() {
	characters := map[string]int{
		"James": 		32,
		"MoneyPenny": 	27,
		"jaws":			42,
	}

	fmt.Println("Before deletion:", characters)

	toDelete := "jaws"
	v, ok := characters[toDelete]
	fmt.Println(v, ok)
	if ok != false {
		delete(characters, toDelete)
	}

	moneypenny := "MoneyPenny"
	if v, ok := characters[moneypenny]; ok {
		fmt.Printf("found %v, deleting %v...\n", moneypenny, v)
		delete(characters, moneypenny)
	}

	fmt.Println("after deleting jaws:", characters)

}

func forRangeOnMap() {
	family := map[string]string{
		"angel":	"male",
		"luz":		"female",
		"julio":	"male",
		"artemis":	"female",
		"jose":		"male",
		"amaris":	"female",
	}

	for k, v := range family {
		fmt.Printf("%v: %v\n", k, v)
	}
}

// creating and adding values as it's created + querying map
func createAndQueryMap() {
	m := map[string]int{
		"James":			32,
		"Miss Moneypenny":	27,
		"jaws":				45,
	}

	fmt.Println(m);
	fmt.Println(m["James"])
	fmt.Println(m["Miss Moneypenny"])

	// when querying a map the return types are the value, which defaults to 0
	// and a boolean,  which lets you know if the value was found or not
	// this is known as the "comma ok" idiom
	v, ok := m["James"]
	fmt.Println(v, ok)

	// common block of code for doing some action if a value is found in a map
	// called the "comma ok" idiom
	if v, ok := m["jaws"]; ok {
		fmt.Println("I was able to find Jaws! Whose age is:", v)
	}
}

// using make func to make a map
func makeMap() {
	m := make(map[string]int)

	m["angel"] = 11
	m["luz"] = 02

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
}


var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}

func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    fmt.Println("unknown time zone:", tz)
    return 0
}