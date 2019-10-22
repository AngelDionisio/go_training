package main

import "fmt"

func main() {
	ch := make(chan int)

	// send
	// this goroutine gets deployed async and goes on its way
	go sendToChannel(ch, 42)

	// because this function is not asynchronous (no go before it)
	// it will be blocking, so this line asserts that the program waits until the sendToChannel
	// finishes and that it is read from it then continues
	getFromChannel(ch)

	// receive
	fmt.Println("exiting program")
}

// send only channel function
func sendToChannel(c chan<- int, value int) {
	c <- value
}

// receive value from channel
func getFromChannel(c <-chan int) {
	fmt.Println(<-c)
}
