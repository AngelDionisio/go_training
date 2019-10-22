package main

import (
	"fmt"
)

func main() {
	addTwo := addToClosure(2)
	addTen := addToClosure(10)

	fmt.Printf("%v\t%v\n", addTwo(10), addTen(10))

	a := counter()
	b := counter()

	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2

	factorialResult := factorial(4)
	fmt.Println(4*3*2*1, factorialResult, loopFactorial(4))
}

// function takes an int
// returns a function that takes an in as argument and returns an int
func addToClosure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

// inner function has a closure on x, which gets initialized with
// int's default value, zero. Every instance of this function has its
// own reference to x
func counter() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// recursive func
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// prefer this as not recursive functions do not have to unwind
func loopFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total = total * n
	}
	return total
}
