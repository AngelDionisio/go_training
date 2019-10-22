package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// the error type in go is just an interface type.
// to implement a custom error all you need is to implement this interface
// type error interface {
//     Error() string
// }
// any type that implements an Error function that returns a string is ALSO of type error

func main() {
	createTextFile("names.txt")
	readFile("names.txt")
}

func createTextFile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // close file before exiting function

	r := strings.NewReader("James Bond")

	io.Copy(f, r)
}

func readFile(fileName string) {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sb, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sb))
}

func fmtNumOfBytesWritten() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}

func askingForInputs() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favorite Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favorite Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
}
