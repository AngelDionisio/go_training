package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/angeldionisio/go_training/channel/worker_pool/pool"
)

const (
	workerCount = 2
	jobCount    = 10
)

func main() {
	log.Println("starting application....")
	collector := pool.StartDispatcher(workerCount)
	comicsToFetch := generateSliceOfRandomNumbers(jobCount)
	log.Printf("Comics: [%v], len: [%d]\n", comicsToFetch, len(comicsToFetch))

	for i, job := range comicsToFetch {
		collector.Work <- pool.Work{ComicID: job, ID: i}
	}
}

func generateSliceOfRandomNumbers(sampleSize int) []int {
	list := make([]int, sampleSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		list[i] = rand.Intn(2300)
	}
	return list
}
