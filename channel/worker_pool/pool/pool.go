package pool

import (
	"fmt"
	"log"

	comic "github.com/angeldionisio/go_training/channel/worker_pool/comic"
)

// Work represents work to be done
type Work struct {
	ID      int
	ComicID int
}

// Worker grabs a waiting job and then does work
type Worker struct {
	ID            int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	JobChannel    chan Work
	End           chan bool
}

// Start creates a worker goroutine
// waits until there is something to process in the Work channel
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// do work
				c, err := comic.GetComic(job.ComicID, w.ID)
				if err != nil {
					fmt.Printf("Error trying to get comic: [%d] due to: %v", job.ComicID, err)
				}
				fmt.Println("Comic:", c)
			case <-w.End:
				return
			}
		}
	}()
}

// Stop kills process
func (w *Worker) Stop() {
	go func() {
		log.Printf("worker [%d] is stopping", w.ID)
		w.End <- true
	}()
}
