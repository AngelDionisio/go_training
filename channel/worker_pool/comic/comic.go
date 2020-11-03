package work

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

// Comic represents XKCD comic JSON data
type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

// generateURL using base XKCD's base URL
func generateURL(comicID int) string {
	return fmt.Sprintf(baseXkcdURL, comicID)
}

// GetComic gets a comic given an ID
func GetComic(comicID int, workerID int) (comic *Comic, err error) {
	url := generateURL(comicID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}

	fmt.Printf("worker [%d] has successfully retrieved comidID: [%d]", workerID, comicID)

	return comic, nil
}
