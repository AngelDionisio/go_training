package main

import (
	"fmt"
	"strconv"
)

func main() {
	list := []string{"foo", "20", "bar", "3"}
	num, i := findLargestPossibleIntInSlice(list)
	fmt.Printf("index: %v with value: %v\n", i, num)

	// usingMakeToCreateSlices()
	// appendToSlice()

}

/* 
1. SLICES in go are a compound type which holds a pointer to an array that holds the values, a size, and maximum capacity
		the make(TYPE, size, capacity) function creates a slice of said type of size: size, and max capacity
2. when using the append function to a slice, if the value being appened overflows the size of the underliying array
		a new array gets created behind the scenes with double the capacity of the original array
*/ 

func appendToSlice() {
	mySlice := make([]int, 10, 12)
	fmt.Printf("value: %v, current size: %v, current capacity %v\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 11, 12)
	fmt.Printf("value: %v, current size: %v, current capacity %v\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 13)
	fmt.Printf("value: %v, current size: %v, current capacity %v\n", mySlice, len(mySlice), cap(mySlice))

}

// findLargestPossibleIntInSlice :: []string -> (int64, int)
func findLargestPossibleIntInSlice(list []string) (int64,  int) {
	var largest int64
	var largestIndex = -1

	for i, _ := range list {
		// attempt to convert string to int64, if err == nil, means successful conversion
		// v, err := strconv.Atoi(list[i])
		v, err := strconv.ParseInt(list[i], 0, 64)

		if err == nil {
			if largest < v {
				largest = v
				largestIndex = i
			}
		}
	}

  // fmt.Println(largest, largestIndex)
	return largest, largestIndex

}

func usingMakeToCreateSlices() {
	// the built in make function takes a type, size, and maximum capacity
	// it populates the array with the zero value for the TYPE of slice
	intSlice := make([]int64, 5, 100) // [0 0 0 0 0]
	strSlice := make([]string, 5, 100) // [    ]
	fmt.Println(intSlice)
	fmt.Println(strSlice)

	// we can also make a slice using a compound data
	// this creates a slice, and populates it with the following values
	otherSlice := []int{1,2,3,4,5}
	fmt.Println(otherSlice)

	// we can make a slice that holds multiple data types with interface
	multiVariadicSlice := []interface{}{"foo", 3.14, 214}
	multiVariadicSlice = append(multiVariadicSlice, true)
	fmt.Printf("type: %T\t length: %v\t capacity: %v\n", multiVariadicSlice, len(multiVariadicSlice), cap(multiVariadicSlice))
	fmt.Println(multiVariadicSlice)
}

func multiDimensionalSlice() {
	jb := []string{"James", "Bond", "super agent"}
	fmt.Println(jb) // [James Bond super agent]
	mp := []string{"Miss", "Moneypenny", "handles the money"}
	fmt.Println(mp) // [Miss Moneypenny handles the money]

	// creating a slice, of a slice of string
	sliceOfPeople := [][]string{jb, mp}
	fmt.Println(sliceOfPeople) // [[James Bond super agent] [Miss Moneypenny handles the money]]
}