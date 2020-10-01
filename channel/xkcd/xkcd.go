package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

// Comic represents XKCD comic JSON data
type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

func main() {
	comicsNeeded := generateSliceOfRandomNumbers(10)
	fmt.Printf("%#v\n", comicsNeeded)

	sequentialFetches(comicsNeeded)
	requestComicsAsync(comicsNeeded)
	comics := getComicsAsync(comicsNeeded)
	fmt.Printf("Comics:  %#v\n", comics)
}

// generateURL using base XKCD's base URL
func generateURL(comicID int) string {
	return fmt.Sprintf(baseXkcdURL, comicID)
}

// getComic gets a comic given an ID
func getComic(comicID int) (comic *Comic, err error) {
	url := generateURL(comicID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}

	return comic, nil
}

// getComicForGoRoutine gets a comic given an ID
// communicates the result via a channel, accepts a pointer to a WaitGroup
// so that it can be coordinated by caller
func getComicForGoRoutine(comicID int, c chan Comic, wg *sync.WaitGroup) error {
	defer wg.Done()

	url := generateURL(comicID)
	resp, err := http.Get(url)
	if err != nil {
		c <- Comic{}
	}

	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		fmt.Printf("Error decoding comicID: %v\n", comicID)
		return err
	}
	c <- comic

	return nil
}

// getComicsAsync exectutes requests async using Goroutines,
// communicates results via a channel.
// coordinates waits using sync.WaitGroup.
func getComicsAsync(listOfIDs []int) map[int]Comic {
	start := time.Now()
	defer func() {
		fmt.Printf("Channels fetches took: %v\n", time.Since(start))
	}()

	var wg sync.WaitGroup
	c := make(chan Comic)

	// wait for all jobs in WaitGroup to complete, then close channel
	go func() {
		wg.Wait()
		close(c)
	}()

	for _, comicID := range listOfIDs {
		wg.Add(1)
		go getComicForGoRoutine(comicID, c, &wg)
	}

	m := make(map[int]Comic)
	for comic := range c {
		m[comic.Num] = comic
	}
	return m
}

func sequentialFetches(comicIDsToFetch []int) {
	start := time.Now()
	defer func() {
		fmt.Println("Sequential Execution Time:", time.Since(start))
	}()

	comicMap := make(map[int]*Comic, len(comicIDsToFetch))

	for _, id := range comicIDsToFetch {
		comic, err := getComic(id)
		if err != nil {
			continue
		}
		comicMap[id] = comic
		fmt.Printf("Fetched comic %d  with title %v\n", id, comic.Title)
	}
}

// requestComicsAsync uses closures and goroutines to send a request for each
// comic async. For each comic we add one go rounine that fetches the request,
// and use a WaitGroup so we can wait for each routine to finish.
// We use an IFFE to closure each request, and it's addition to the result map.
// It then waits for all Goroutines to complete before exiting.
func requestComicsAsync(listOfComicIDs []int) map[int]*Comic {
	var wg sync.WaitGroup
	comicMap := make(map[int]*Comic)

	for _, comicID := range listOfComicIDs {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // signal goroutine completion before exiting closure
			comic, err := getComic(id)
			if err != nil {
				return
			}

			fmt.Printf("Fetched comicID: %v, with title: %v\n", id, comic.Title)
			comicMap[id] = comic
		}(comicID)
	}

	wg.Wait()

	return comicMap
}

func generateSliceOfRandomNumbers(size int) []int {
	list := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(2000)
	}
	return list
}
