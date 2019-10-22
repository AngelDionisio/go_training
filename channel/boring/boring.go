package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joeChan := boring("Joe")
	annChan := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joeChan)
		fmt.Println(<-annChan )
	}

	fmt.Println("You are both boring; I'm leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}