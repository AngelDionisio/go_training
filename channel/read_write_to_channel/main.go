package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	// create bidirectional channel, that accepts ints
	c := make(chan int)

	// create goroutine, which starts adding values to channel
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		// safe to close because this channel gets returned while values are still being put into the channel?
		// close is used here to signal that the channel will no longer accept any more values. Which in turn,
		// let's the for-range in the recieve func to know when to end.
		// a closed channel can be read from, but it will return the zero value for its type.
		close(c)
	}()

	// channel gets returned before closing, as values keep being added to it
	return c
}

func receive(c <-chan int) {
	// receive starts pulling messages from channel
	// the reading from channel is blocking, which forces the meeting with the other routine to pass the baton
	for v := range c {
		fmt.Println(v)
	}
}
