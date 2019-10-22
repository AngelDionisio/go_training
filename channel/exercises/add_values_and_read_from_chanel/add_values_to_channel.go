package main

import (
	"fmt"
)

// write a program that adds 100 numbers to a channel
// pull the numbers from the channel and read them

func addValuesToChannel(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func readFromChannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)

	go addValuesToChannel(ch)

	readFromChannel(ch)

	fmt.Println("about to exit")
}
