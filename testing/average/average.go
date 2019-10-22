package main

import (
	"fmt"
)

func main() {
	nums := []float64{1.5, 40.2, 15.35}
	result := Average(nums)
	fmt.Println(result)
}

// Average takes a slice of numbers and returns its average
func Average(ls []float64) float64 {
	var sum float64
	var avg float64

	for _, v := range ls {
		sum = sum + v
	}

	avg = sum / float64(len(ls))

	return avg
}
