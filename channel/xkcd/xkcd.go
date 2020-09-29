package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Comic represents XKCD comic JSON data
type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

func getComic(comicID int) (comic *Comic, err error) {
	url := fmt.Sprintf(baseXkcdURL, comicID)
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

func main() {
	// start := time.Now()
	// defer func() {
	// 	fmt.Println("Execution Time:", time.Since(start))
	// }()

	comicsNeeded := generateSliceOfRandomNumbers(10)
	sequentialFetches(comicsNeeded)
	waitGroupFetches(comicsNeeded)
}

func sequentialFetches(comidIDsToFetch []int) {
	start := time.Now()
	defer func() {
		fmt.Println("Sequential Execution Time:", time.Since(start))
	}()

	comicMap := make(map[int]*Comic, len(comidIDsToFetch))

	for _, id := range comidIDsToFetch {
		comic, err := getComic(id)
		if err != nil {
			continue
		}
		comicMap[id] = comic
		fmt.Printf("Fetched comid %d  with title %v\n", id, comic.Title)
	}
}

func waitGroupFetches(comicIdsToFetch []int) map[int]*Comic {
	start := time.Now()
	defer func() {
		fmt.Println("waitGroup Execution Time:", time.Since(start))
	}()

	comicMap := make(map[int]*Comic, len(comicIdsToFetch))
	wg := sync.WaitGroup{}

	// set number of goRoutines to wait for
	wg.Add(len(comicIdsToFetch))
	for _, id := range comicIdsToFetch {
		// wg.Add(1)
		go func(id int) {
			defer wg.Done()
			comic, err := getComic(id)
			if err != nil {
				return
			}

			comicMap[id] = comic
			fmt.Printf("Fetched comid %d  with title %v\n", id, comic.Title)
			// wg.Done()

		}(id)
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
