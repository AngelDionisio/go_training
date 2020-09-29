package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	checkLinks(links)
}

func checkLinks(URLs []string) {
	var wg sync.WaitGroup
	c := make(chan string)

	// launch a Goroutine that will block until all requests have been completed, then close channel
	go func() {
		wg.Wait()
		close(c)
	}()

	for _, url := range URLs {
		wg.Add(1) // tell WaitGroup that there is another item to wait for
		go checkLink(url, c, &wg)
	}

	for msg := range c {
		fmt.Println(msg)
	}

}

func checkLink(link string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down!"
		return
	}

	c <- link + " is up!"
}
