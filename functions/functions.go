package main

import (
	"fmt"
)

// func (r receiver) identifier(parameters) (return(s)) { ... }
// parameters: what a function recieves
// arguments: what  you are passing to the function
// every thing in Go is PASS BY VALUE
// functions can return more than one value e.g. func foo() (string int)

func main() {
	// sum := add(1, 2, 3, 4, 5)
	// fmt.Println("the total is:", sum)

	// unfurling
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,}
	result := add(nums...)
	fmt.Println("unfolding result:", result)

	// fmt.Println("zero values:", add())

	// names := []string{"joe", "anna", "trevor"}

	// greeting("hello:", names...)

	// defer foo()
	// bar()

	// anonymousFunctions()
	functionExpression()

	holdsFunction := returnFunction()
	fmt.Printf("%T\n", holdsFunction) // funct() int
	x := holdsFunction()
	fmt.Println(x) // 451

	sumFn := func(nums ...int) int {
		var sum int
		for _, v := range nums {
			sum += v
		}
		return sum
	}

	fmt.Println("adding ALL nums", sumFn(nums...))

	a := addEven(sumFn, nums...)
	fmt.Println("result of add even:", a)

}

// callback, passing a function as an argument to another function
// first param is a function,
// said function takes unlimited number of ints and returns an int
// second param is an unlimited number of ints
// the result of addEven is an int
func addEven(f func(xi ...int) int, nums ...int) int {
	var evenNums []int
	for _, v := range nums {
		if v % 2 == 0 {
			evenNums = append(evenNums, v)
		}
	}
	fmt.Println("slice of even numbers:", evenNums)

	// now let's pass the slice of even to f and return result
	return f(evenNums...)
}

// function that returns a function
// here we are declaring a function that returns
// a function of type func() int, which in turn,
// will return 451
func returnFunction() func() int {
	return func() int {
		return 451
	}
}

func functionExpression() {
	add := func(x int, y int) int { 
		return x + y;
	}

	fmt.Println("function expression executed with result:", add(2,3))
}

// anonymous self executing functions
func anonymousFunctions() {
	// anonymous (self executing function)
	func (x int) {
		fmt.Println("anonymous function ran", x)
	}(42)
}

// variadic parameters (zero or more values)
// if zero params are present the value passed is nil
// the type of ...TYPE will be a slice of TYPE ([]type)
// if function accepts multiple params, the variadic one has to be the last one
// If function is variadic with final parameter p of type ...T, then  within function
//  the type of p is equivalent to []T (a slice of T). If f is invoked with no actual
//  parameters for p, then the value passed to p is nil.
// this function takes an unlimited number of ints
func add(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum

}

func greeting(prefix string, who ...string) {
	fmt.Println(prefix, who)
}

// A "defer" statement invokes a function whose execution is deferred to the moment the surrounding
// function returns, either because the surrounding function executed a return statment,
// reached the end of it's function body, or because the corresponding goroutine is panicking
// basically an after, local to the surrounding function
// an example of good use: if you are writing a function that opens a file,
// right at the top, we can defer the function that closes the file, ensuring it releases it.

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}