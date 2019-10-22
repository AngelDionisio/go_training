package main

import (
	"fmt"
	"runtime"
	"sync"
)

// A "go" statement starts the execution of a function call as an independent concurrent thread of control
// or goroutine, within the same address space.
// The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions
// are restricted fas for expression statements.

// var wg of type sync.WaitGroup
var wg sync.WaitGroup

func main() {
	fmt.Println("************** START ******************")

	wg.Add(1)

	go foo()
	bar()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("************** END ******************")

}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar", i)
	}
}


// This example fetches several URLs concurrently,
// using a WaitGroup to block until all the fetches are complete.

/*

var wg sync.WaitGroup
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
wg.Wait()

*/