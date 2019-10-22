package main

import (
	"fmt"
)

type user struct {
	name   string
	gender string
}

type categorized struct {
	male   []string
	gender []string
}

func main() {
	var categorized = make(map[string][]string)

	users := []user{
		{name: "Raphel", gender: "male"},
		{name: "Tom", gender: "male"},
		{name: "Jerry", gender: "male"},
		{name: "Dorry", gender: "female"},
		{name: "Suzie", gender: "female"},
		{name: "Dianna", gender: "female"},
		{name: "Prem", gender: "male"},
	}

	for _, v := range users {
		if v.gender == "male" {
			categorized["male"] = append(categorized["male"], v.name)
		}
		if v.gender == "female" {
			categorized["female"] = append(categorized["female"], v.name)
		}
	}

	fmt.Println(categorized)
	fmt.Println("about to exit")
}
