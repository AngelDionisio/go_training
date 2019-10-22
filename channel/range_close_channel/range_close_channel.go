package main

import "fmt"

func main() {
	ch := make(chan int)

	// send
	go insertMultipleValuesToChannel(ch)

	// recieve
	// when you range over a channel, it will keep printing values until the channel is closed
	for v := range ch {
		fmt.Println(v)
	}
}

func insertMultipleValuesToChannel(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	// if we do not close the channel, the channel would be waiting for more values to come,
	// so we would get an "all goroutines are sleep" error
	close(c)

}
