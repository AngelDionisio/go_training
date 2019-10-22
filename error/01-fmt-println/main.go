package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error happened:", err)
		// log.Println("err happened:", err)
		// log.Fatalln(err)
		// panic(err)
	}
}
