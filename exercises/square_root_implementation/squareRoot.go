package main

import (
	"fmt"
)

// Sqrt calculates the square root of a positive number
func Sqrt(x float64) float64 {
	maxIterations := 50
	z := float64(1)
	lastVal := float64(-99999)
	currentIter := 0

	for i := 0; i < maxIterations; i++ {
		z -= (z*z - x) / (2 * z)
		if z == lastVal {
			return z
		}
		lastVal = z
		currentIter++
		fmt.Println(
			"current iteration:", currentIter,
			"value:", z,
			"lastValue:", lastVal)
	}

	fmt.Println("it took longer than 100 iterations, returning current approximation:", z)
	return z
}

func main() {
	fmt.Println(Sqrt(1048576))
}
