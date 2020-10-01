package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	checkURLs(urls)
}

func checkURL(url string) string {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v is down!\n", url)
		return fmt.Sprintf("%v is down!\n", url)
	}

	return fmt.Sprintf("%v is up!", url)
}

func checkURLs(urlList []string) {
	var wg sync.WaitGroup
	c := make(chan string)

	for _, url := range urlList {
		wg.Add(1)
		go func(lnk string) {
			defer wg.Done()
			c <- checkURL(lnk)
		}(url)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}
