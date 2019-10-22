package main

import "fmt"

/*
	The recover built-in function allows a program to manage behavior of a panicking goroutine.
	Recover can ONLY be called from a defer function.
	Executing a recover inside a deferred function (but not any function called by it) stops the panicking sequence
	by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside
	deferred function it will not stop the panicking sequence.
	https://blog.golang.org/defer-panic-and-recover
*/

/*
	A DEFER statement pushes a function call to a list. The list of saved calls is executed after the surrounding function returns.
	Defer is commonly used to simplify functions that return various clean-up actions, like closing files.
	1. A deferred function's argument are evaluated when the defer statement is evaluated. (at the time it was called)
			func a() {
    			i := 0
    			defer fmt.Println(i)
    			i++
    			return
			} // this defer function will print 0

	2. Deferred function calls are executed in Last in, First out order order after the surrounding function returns.
			func b() {
    			for i := 0; i < 4; i++ {
        		defer fmt.Print(i)
    			}
			} // this function will print 3, 2, 1, 0

	3. Deferred functions may read and assing to returning function's named return values.
			func c() (i int) {
    			defer func() { i++ }()
    			return 1
			} // this will return 2
*/

/*
	PANIC is a built-in function that stops the ordinary flow of control and begins `panicking`. When the function F calls panic,
	execution of F stops, any deferred functions are executed normally, and F returns to its caller. To the caller, F then behaves like a call to panic.
	The process continues up the stack until all functions in the current goroutine have returned, at this point, the program crashes. Panics can be initiated
	by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

	RECOVER is a built-in function that regains control of a panicking goroutine. Recover is only useful INSIDE deferred functions. During normal execution,
	a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic
	and resume normal execution.

*/

func main() {
	f()
	fmt.Println("normally recovered from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g") // this line never gets executed, on a panic only deferred functions exec
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer function in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
