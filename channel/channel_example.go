package main

import (
	"fmt"
)

// channels DISCARD values returned by functions
// channels BLOCK. Channels need to run concurrenty with goroutines
// they are like two relay racers trying to pass a baton, both need to interlock
// to then pass the baton, and they take off.

func doSomething(x int) int {
	// does something already built
	return x * 5
}

func bufferedChannel() {
	// this is a buffered channel that can hold two values
	ch := make(chan string, 2)
	go func() {
		ch <- "first msg put in channel"
		ch <- "second msg put in channel"
	}()

	fmt.Println(<-ch) // 42
	fmt.Println(<-ch) // 43
}

// a buffered channel allows you to put values in a channel
// regardless if there is something to pull it off (e.g. Goroutine)
func successfullBufferChannelBlock() {
	ch := make(chan string, 1)

	// because channels block, 
	ch <- "buffers block!"

	fmt.Println(<-ch)
}

func main() {
	// unbuffered channel, a buffered channel has an extra param as an integer
	// which denotes how many messages it can hold. It's recommended to use unbuffered channels.
	// this "passing of the baton" have to ocurr by two processes that are launched
	// and running asynchronously. Those two async processes are the channel and the goroutine.
	// You can only pull a value "passing of the baton" from a goroutine, using a channel. 
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)

	bufferedChannel()
	successfullBufferChannelBlock()

}

