package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run -race <file_name>.go
// will allow you see if there is any race conditions on your code

func raceCondition() {
	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())


	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// sync.Mutex allows to checkout / lock a piece of data / functionality for reading and writing
	// so that if other Goroutines attempt to access it / modify it, they have to wait
	// until they have been released by any Goroutine that currently has it locked.
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// Gosched yields the processor, allowing other goroutines to run.
			// It does not suspend the current goroutine, so execution resumes automatically.
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			// this must be in the go anonymous self executing function
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	// don't exit program until all goroutines have been marked Done
	wg.Wait()

	fmt.Println("count:", counter)

}

func main() {
	raceCondition()
}

