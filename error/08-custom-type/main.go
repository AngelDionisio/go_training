package main

import (
	"fmt"
	"log"
	"math"
)

type customErrorWithLocation struct {
	lat  string
	long string
	err  error
}

// norgateMathError now implements the Error interface, now making also of type Error
/*
	type error interface {
        	Error() string
	}
*/
func (n customErrorWithLocation) Error() string {
	return fmt.Sprintf("a norgate math error ocurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	v, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(v)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		customErr := fmt.Errorf("norgate math error: attempted to square root of a negative number: %v", f)
		return 0, customErrorWithLocation{"40.736343 N", "-73.993642 W", customErr}
	}

	return math.Sqrt(f), nil
}
