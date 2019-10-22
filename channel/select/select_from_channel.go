package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go populate(even, odd, quit)

	retreive(even, odd, quit)

	fmt.Println("about to exit")
}

func populate(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func retreive(e, o, q <-chan int) {
	// infinite loop
	for {
		// select based on channel and type
		select {
		case v := <-e:
			fmt.Println("from even channel:\t", v)

		case v := <-o:
			fmt.Println("from odd channel:\t", v)

		case v := <-q:
			fmt.Println("from quit channel:\t", v)
			return
		}
	}
}
