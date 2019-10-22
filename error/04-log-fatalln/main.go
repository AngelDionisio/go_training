package main

import (
	"fmt"
	"log"
	"os"
)

// exits deferred functions DO NOT RUN
// equivalent to running Println() followed by a call to os.Exit(1)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("I am running from foo!")
}
