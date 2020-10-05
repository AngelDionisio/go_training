package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
	"time"
)

// ToDo represents a to do from "https://jsonplaceholder.typicode.com/todos/"
type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// IsEmpty checks for an empty ToDo
func (t ToDo) IsEmpty() bool {
	return reflect.DeepEqual(ToDo{}, t)
}

// CustomResponse represents a response with its ToDo and its correspondingID
type CustomResponse struct {
	ID    int
	TODO  ToDo
	Error string
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
		"https://jsonplaceholder.typicode.com/todos/-1",
	}
	checkURLs(urls)
}

func checkURLs(urls []string) {
	c := make(chan *CustomResponse)
	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1) // adding pending operation to waitgroup
		go checkURL(link, c, &wg)
	}

	// creating an IFFE Goroutine to handle waiting for all requests to complete
	// and closing the channel
	go func() {
		wg.Wait() // blocks goroutine until the WaitGroup counter is zero
		close(c)
	}()

	var responses []CustomResponse
	for msg := range c {
		responses = append(responses, *msg)
		fmt.Printf("Processed response: %#v\n", msg)
	}

}

func checkURL(url string, c chan *CustomResponse, wg *sync.WaitGroup) {
	// defer (*wg).Done()
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		errResponse := &CustomResponse{
			Error: "We could not reach URL:" + url,
		}
		c <- errResponse
	}

	toDoItem, err := unmarshalToDo(res)
	if err != nil {
		return
	}

	c <- toDoItem
}

func readResponseBody(res *http.Response) ([]byte, error) {
	// close the response once we return from the function
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response due to:", err)
		return nil, err
	}
	return content, nil
}

func unmarshalToDo(res *http.Response) (*CustomResponse, error) {
	var toDoResp ToDo
	content, err := readResponseBody(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &toDoResp)
	if err != nil {
		fmt.Println("Could not unmarshal due to:", err)
		return nil, err
	}

	if toDoResp.IsEmpty() {
		emptyResponseError := &CustomResponse{
			ID:    -1, // pass actual ID
			Error: "Empty response",
		}
		return emptyResponseError, nil
	}

	resp := &CustomResponse{
		ID:   toDoResp.ID,
		TODO: toDoResp,
	}
	return resp, nil
}
