package main

import (
	"fmt"
	"log"
	"os"
)

// exits deferred functions DO NOT RUN
// equivalent to running Println() followed by a call to panic()

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panic(err)
		// panic() // will run defer functions and run throught he call stack running defer functions as it goes
	}
}

func foo() {
	fmt.Println("I am running from foo!")
}
