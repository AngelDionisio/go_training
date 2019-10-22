package main

import (
	"fmt"
	"sort"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type bySayings []person

// bySayings implements the sort.Interface for []person based on the Sayings field
// slice of person is also of type bySayings, which allows us to use the Len() func on slice of ints
// like so bySayings.Len(users)
// in order to use sort, we need to implent the sort.Interface, which needs a Len, Swap and Less
// Less is the function provided to sort.Sort by default, you can also pass a default sort function to Sort
func (bs bySayings) Len() int           { return len(bs) }
func (bs bySayings) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs bySayings) Less(i, j int) bool { return bs[i].Sayings[i] < bs[j].Sayings[j] }

// sort the []user by age, last
// also sort each []string "Sayings" for each user
func sortStrucValues() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "A",
		Last:  "Zzzz",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	for _, p := range users {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Printf("\t%v\n", s)
		}

	}

	// sort by FIRST name
	sort.Slice(users, func(i, j int) bool {
		return users[i].First < users[j].First
	})

	fmt.Println("***************** AFTER SORT BY FIRST NAME AND SORT INNER SAYINGS *****************")
	for _, p := range users {
		fmt.Println(p.First, p.Last, p.Age)
		// sort Sayings []string alphabetically
		sort.Strings(p.Sayings)
		for _, s := range p.Sayings {
			fmt.Printf("\t%v\n", s)
		}

	}

	// sort by LAST name
	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})

	fmt.Println("***************** AFTER SORT BY LAST NAME *****************")
	for _, p := range users {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Printf("\t%v\n", s)
		}

	}

	// sort by provided bySavings Less function
	sort.Sort(bySayings(users))

	fmt.Println("***************** SORTING EACH SAYING *****************")
	for _, p := range users {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Printf("\t%v\n", s)
		}

	}

	// we can also sort by a custom Less function
	sort.Slice(users, func(i, j int) bool {
		return users[i].Sayings[i] > users[j].Sayings[j]
	})

	fmt.Println("***************** AFTER SORT DECREASING *****************")
	for _, p := range users {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Printf("\t%v\n", s)
		}

	}

	fmt.Println(bySayings.Len(users))
}

func sortingIntsAndStrings() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

// Marshal
// func Marshal(v interface{}) ([]byte, error)
// returns the JSON encoding of v

// Unmarshal
// func Unmarshal(data []byte. v interface{}) error
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
// helpful site: https://mholt.github.io/json-to-go/

func main() {
	sortStrucValues()
	// sortingIntsAndStrings()
}
