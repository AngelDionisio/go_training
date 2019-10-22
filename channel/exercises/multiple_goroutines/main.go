package main

import "fmt"

func main() {
	const numberOfRoutines = 10
	ch := make(chan int)

	for i := 0; i < numberOfRoutines; i++ {
		go addValuesToChannel(10, ch)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)
	}

	// for i := 0; i < numberOfRoutines; i++ {
	// 	go func() {
	// 		for n := 0; n < numberOfRoutines; n++ {
	// 			ch <- n
	// 		}
	// 	}()
	// }

	// you can for range on a channel that is already closed.
	// it is possible to close a non empty channel, but still have the remaining values be received
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	fmt.Println("about to exit")
}

func addValuesToChannel(numOfValues int, c chan int) {
	for i := 0; i < numOfValues; i++ {
		c <- i
	}
}
