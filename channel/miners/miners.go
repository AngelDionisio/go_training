package main

import (
	"fmt"
	"time"
)

// the main function runs on its own goroutine.
// once the main function returns, it closes all other go routines that are currently running.
// good practice is to create a "done channel", which the main function blocks on waiting to read.
// once we finish our  work we write to this channel, and the program will end.
func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	// unbuffered channels take only one message at a time.
	// sender blocks until there is another goroutine that reads the value (taking it from the channel)
	oreChannel := make(chan string)
	minedOreChannel := make(chan string)
	quitChan := make(chan string)

	go func() {
		quitChan <- "that's all folks!"
	}()

	// Finder
	go func(mine []string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(theMine)

	// Miner / Ore Breaker
	go func() {
		for foundOre := range oreChannel {
			fmt.Println("Miner: received", foundOre, "from finder")
			minedOreChannel <- foundOre
		}
		// Note: Ranging over a channel will block until another item is sent to the channel.
		// The only way to stop the goroutine from blocking after all sends have ocurred is
		// by closing the channel with 'close(channel)'

		// for i := 0; i < 3; i++ {
		// 	foundOre := <-oreChannel // read from oreChannel
		// 	fmt.Println("Miner: received", foundOre, "from finder")
		// 	minedOreChannel <- foundOre
		// }
	}()

	// Smelter
	go func() {
		for minedOre := range minedOreChannel {
			fmt.Println("Received from Miner:", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	// <-quitChan
	<-time.After(time.Second * 5) // Again, ignore this for now

	// fmt.Println("From finder:", finder(theMine))
	// fmt.Println("From miner:", miner(finder(theMine)))
	// fmt.Println("From smelter:", smelter(miner(finder(theMine))))

	// bufferedChannels()

	fmt.Println("about to exit")
}

func bufferedChannels() {
	// a buffered channel can hold "int" number of data before needing
	// another goroutine to read from it
	bufferedChan := make(chan string, 3)

	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()

	<-time.After(time.Second * 1)

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
	}()

	<-time.After(time.Second * 3)
}

func finder1(ls []string) []string {
	var ore []string

	for _, v := range ls {
		if v == "ore" {
			ore = append(ore, v)
			fmt.Println("finder 1: found ore")
		}
	}
	return ore
}

func finder2(ls []string) []string {
	var ore []string

	for _, v := range ls {
		if v == "ore" {
			ore = append(ore, v)
			fmt.Println("finder 2: found ore")
		}
	}
	return ore
}

func miner(ls []string) []string {
	var minedOre []string

	for range ls {
		minedOre = append(minedOre, "minedOre")
	}

	return minedOre
}

func smelter(ls []string) []string {
	var smeltedOre []string

	for range ls {
		smeltedOre = append(smeltedOre, "smeltedOre")
	}

	return smeltedOre
}
