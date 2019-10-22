package main

import "fmt"

// channels DISCARD values returned by functions
// channels BLOCK. Channels need to run concurrenty with goroutines
// they are like two relay racers trying to pass a baton, both need to interlock
// to then pass the baton, and they take off.

func main() {
	// receive only channel
	c := make(chan int)
	ch := make(chan<- string)

	go func() {
		ch <- "value inteserted into Channel inside Goroutine"
	}()

	// this will not work, the channel is a receieve only channel
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("ch\t%T\n", ch)

}
