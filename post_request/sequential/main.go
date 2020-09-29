package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ToDo represents a to do from "https://jsonplaceholder.typicode.com/todos/"
type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()

	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/6",
	}
	checkURLs(urls)
}

func checkURLs(urls []string) {
	for _, link := range urls {
		checkURL(link)
	}
}

func checkURL(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Was not able to reach: %v\n", url)
		return
	}

	fmt.Println("Success reaching URL:", url)

	unmarshalToDo(res)
}

func readResponseBody(res *http.Response) []byte {
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response due to:", err)
		return nil
	}
	return content
}

func unmarshalToDo(res *http.Response) {
	var toDO ToDo
	content := readResponseBody(res)
	err := json.Unmarshal(content, &toDO)
	if err != nil {
		fmt.Println("Could not unmarshal due to:", err)
	}
	fmt.Printf("Unmarshaled Object: %#v\n", toDO)
}
