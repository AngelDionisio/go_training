package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	v, err := sqrtFormatPrinting(-10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(v)
}

func sqrtErrorsNew(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("cannot take square root of negative numbers")
	}
	return math.Sqrt(f), nil
}

func sqrtFormatPrinting(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: cannot take sqrt of negrative number: %v", f)
	}
	return math.Sqrt(f), nil
}
