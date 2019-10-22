package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person

// byAge implements the sort.Interface, which requires these three functions
// https://godoc.org/sort#Interface
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byname []person

func (bn byname) Len() int           { return len(bn) }
func (bn byname) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byname) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {
	people := []person{
		{"Q", 42},
		{"Moneypenny", 27},
		{"M", 52},
		{"James", 36},
	}

	fmt.Println(people)
	// using type conversion from people, we can use sort.Sort on a type
	// that implements the sort.Interface. Using the Less as default
	sort.Sort(byAge(people))
	fmt.Println("after sorting by age:", people)
	// using the length function implemented byAge
	fmt.Println("length of people object:", byAge(people).Len())

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with byAge.Less.
	// using sort.SliceStable as go documentation suggests this one
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("after sorting by custom implementation of Less:", people)

	sort.Sort(byname(people))
	fmt.Println("after being sorted by name:", people)
}
