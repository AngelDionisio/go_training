package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// fan-out is the concept of sending some processes that are independent of each other
// to multiple goroutines so that they can happen at the same time. Things like, encoding 100 videos.
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// went from general (bi-directional channel), to specific (send channels - channels that receive values only)
func populate(c chan<- int) {
	for i := 0; i < 50; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		// this is fanning out, for each value in c1, we are launching a Goroutine doing some time consuming work
		// is this how we should launch async requests?
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	// wait for all waitgroups added to complete
	wg.Wait()
	close(c2)
}

func timeConsumingWork(x int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return x + rand.Intn(1000)
}
