package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go populate(even, odd, quit)

	retreive(even, odd, quit)
	// receive(even, odd, quit)

	fmt.Println("about to exit")
}

func populate(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	close(quit)
}

func retreive(even, odd <-chan int, quit <-chan bool) {
	// infinite loop
	for {
		// select based on channel and type
		select {
		case v := <-even:
			fmt.Println("from even channel:\t", v)
		case v := <-odd:
			fmt.Println("from odd channel:\t", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma, ok:\t\t", i, ok)
				return
			}
		}
	}
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:\t", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:\t", v)
		case _, _ = <-quit:
			return
		}
	}
}
