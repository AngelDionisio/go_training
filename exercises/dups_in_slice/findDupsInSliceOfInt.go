package main

import (
	"fmt"
	"reflect"
)

// reflect.DeepEqual
func main() {
	nums := []int{1, 2, 5, 3, 4, 5, 2}
	dupes := findDupsInSlice(nums)

	fmt.Println(dupes)
}

func findDupsInSlice(ls []int) []int {
	var seen, dups []int
	seen2 := reflect.TypeOf(ls)
	fmt.Println("reflect type of:", seen2)

	for _, v := range ls {
		if !containsInt(seen, v) {
			seen = append(seen, v)
		} else {
			dups = append(dups, v)
		}
	}

	return dups
}

func contains(ls []interface{}, e interface{}) bool {
	for _, v := range ls {
		if v == e {
			return true
		}
	}

	return false
}

func containsInt(ls []int, e int) bool {
	for _, v := range ls {
		if v == e {
			return true
		}
	}

	return false
}
