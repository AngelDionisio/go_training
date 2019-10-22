package main

import (
	"fmt"
	"log"
	"os"
)

// standar logger defaults to fmt.Println

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("error happened:", err)
		log.Println("err happened:", err)
		// log.Fatalln(err)
		// panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in same directory")
}
