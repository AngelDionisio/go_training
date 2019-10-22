package main

import (
	"fmt"
	"testing"
)

/*
Make code ready to BET on (Benchmark, Example and Test)
Tests must:
	be in a file that ends with "_test.go"
	put the file in the same package as the one being tested
	be in a fucn with a signature "func TestXxxx(*testing.T)" takes a pointer of type testing.T

Run a test
	go test

Deal with test failure
	use t.Error to singal failure

Write results to file
	go test -coverprofile filename.extension
	go tool cover -html=filename (creates an HTML report from test coverage)
*/

// TestAverage tests multiple test cases using a table of tests
func TestAverage(t *testing.T) {
	type test struct {
		data   []float64
		answer float64
	}

	tests := []test{
		test{[]float64{90, 85, 79, 86}, 85},
		test{[]float64{1, 2}, 1.5},
		test{[]float64{-1, 2, 5}, 2},
		test{[]float64{1.5, 40.2, 15.35}, 19.01666666666667},
	}

	for _, v := range tests {
		// var v float64
		avg := Average(v.data)
		if avg != v.answer {
			t.Error("expected", v.answer, "got ", avg)
		}
	}

}

// to make Examples, naming conventions are defined as 'Example<funtionName>'
func ExampleAverage() {
	var f float64
	f = Average([]float64{1, 2})
	fmt.Println(f)
	// Output:
	// 1.5
}

// naming convetions for benchmarks
// Benchmark<functionName>
func BenchmarkAverage(b *testing.B) {
	xf := []float64{1, 2}
	for i := 0; i < b.N; i++ {
		Average(xf)
	}
}

// TestAverageSingleValue is a function that tests the Average function
// with one single test case
func TestAverageSingleValue(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("expected 1.5, got ", v)
	}
}
